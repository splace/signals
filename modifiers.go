package signals

import "encoding/gob"

func init() {
	gob.Register(Shifted{})
	gob.Register(Spedup{})
	gob.Register(SpedupProgressive{})
	gob.Register(Looped{})
	gob.Register(Inverted{})
	gob.Register(Reversed{})
	gob.Register(Reflected{})
	gob.Register(Modulated{})
	gob.Register(Triggered{})
	gob.Register(Segmented{})
}

// a Function that shifts the x of another function
type Shifted struct {
	Function
	Shift x
}

func (s Shifted) call(t x) y {
	return s.Function.call(t - s.Shift)
}

// a Function that scales the x of another function
type Spedup struct {
	Function
	Factor float32
}

func (s Spedup) call(t x) y {
	return s.Function.call(x(float32(t) * s.Factor))
}

/*
// TODO spedup tone should have MaxX and period changed
// a Function that scales the x of another function
type Squeeze struct {
	LimitedFunction
	Factor float32
}

func (s Squeeze) Call(t x) y {
	return s.LimitedFunction.Call(x(float32(t) * s.Factor))
}

func (s Squeeze) MaxX() x {
	return x(float32(s.LimitedFunction.MaxX())*s.Factor)
}

*/

type SpedupProgressive struct {
	Function
	Rate x
}

func (s SpedupProgressive) call(t x) y {
	return s.Function.call(t + t*t/s.Rate)
}

// a Function that repeats another function
type Looped struct {
	Function
	Length x
}

func (s Looped) call(t x) y {
	return s.Function.call(t % s.Length)
}

func (s Looped) Period() x {
	return s.Length
}

// a Function that produceds y values that are the negative of another functions y values
type Inverted struct {
	Function
}

func (s Inverted) call(t x) y {
	return -s.Function.call(t)
}

// a Function that returns y's that are for the -ve x of another function
type Reversed struct {
	Function
}

func (s Reversed) y(t x) y {
	return s.Function.call(-t)
}

// a Function that produces values that are flipped over, (Maxy<->zero) of another function
type Reflected struct {
	Function
}

func (s Reflected) call(t x) y {
	if r := s.Function.call(t); r < 0 {
		return -maxY - r
	} else {
		return maxY - r
	}
}

// a Function that stretches the x values of another function, in proportion to the value of a modulation function
type Modulated struct {
	Function
	Modulation Function
	Factor     x
}

func (s Modulated) call(t x) y {
	return s.Function.call(t + MultiplyX(float64(s.Modulation.call(t))/maxyfloat64, s.Factor))
}

// TODO if a modulation function is periodic then the modulated will be, or, smaller of either?
/*
// a Function that stretches the time of another function, in proportion to the value of a modulation function
type Modulated struct {
	Function
	Modulation Function
	Factor     float64
}

func (s Modulated) y(t x) y {
	return s.Function.y(x(float64(t) * DB(float64(s.Modulation.y(t))/Maxyfloat64*s.Factor))
}
// a Function that has equal width uniform gradients as an approximation to another function
type Segmented struct {
	Function
	Width   x
}

func (s Segmented) y(t x) y {
	temp:=t%s.Width
	return s.Function.y(t-temp)/y(s.Width)*y(s.Width-temp)+s.Function.y(t+s.Width-temp)/y(s.Width)*y(temp)
}
*/

// a Function that has equal width uniform gradients as an approximation to another function.
// (sebsequent calls within the same segment, are generated from cached end values, so doesn't call embedded Function)
type Segmented struct {
	Function
	Width  x
	i1, i2 x
	l1, l2 y
}

func NewSegmented(s Function, w x) Segmented {
	return Segmented{Function: s, Width: w}
}

func (s Segmented) call(t x) y {
	temp := t % s.Width
	if t-temp != s.i1 || t-temp+s.Width != s.i2 {
		s.i1 = t - temp
		s.i2 = t - temp + s.Width
		s.l1 = s.Function.call(s.i1)
		s.l2 = s.Function.call(s.i2)
	}
	return s.l1/y(s.Width)*y(s.Width-temp) + s.l2/y(s.Width)*y(temp)
}

// Triggered shifts a Function's x to make it cross a trigger y at zero x.
// searches with a Resolution, from Shift+Resolution to MaxShift, then from 0 to Shift.
// Delay is set to last found trigger, so subsequent uses finds new crossing, and wraps round.
// Rising can be alternated to find either way crossing
type Triggered struct {
	Function
	Trigger        y
	Rising         bool
	Resolution     x
	MaxShift       x
	Shift          x
	searched       Function
	locatedTrigger y
	locatedRising  bool
}

func (s *Triggered) call(t x) y {
	if s.Trigger != s.locatedTrigger || s.searched != s.Function || s.locatedRising != s.Rising {
		s.searched = s.Function
		s.locatedTrigger = s.Trigger
		s.locatedRising = s.Rising
		if s.Rising && s.Function.call(s.Shift) > s.Trigger || !s.Rising && s.Function.call(s.Shift) < s.Trigger {
			s.Shift += s.Resolution
		}
		for t := s.Shift; t <= s.MaxShift; t += s.Resolution {
			if s.Rising && s.Function.call(t) > s.Trigger || !s.Rising && s.Function.call(t) < s.Trigger {
				s.Shift = t
				return s.Function.call(t)
			}
		}
		for t := x(0); t < s.Shift; t += s.Resolution {
			if s.Rising && s.Function.call(t) > s.Trigger || !s.Rising && s.Function.call(t) < s.Trigger {
				s.Shift = t
				return s.Function.call(t)
			}
		}
		s.Shift = 0
	}
	return s.Function.call(t + s.Shift)
}
