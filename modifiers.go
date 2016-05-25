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
	gob.Register(RateModulated{})
	gob.Register(Triggered{})
	gob.Register(Segmented{})
}

// a Signal that shifts the x of another Signal
type Shifted struct {
	Signal
	Shift x
}

func (s Shifted) property(t x) y {
	return s.Signal.property(t - s.Shift)
}

// a Signal that scales the x of another Signal
type Spedup struct {
	Signal
	Factor float32
}

func (s Spedup) property(t x) y {
	return s.Signal.property(x(float32(t) * s.Factor))
}

/*
// TODO spedup tone should have MaxX and period changed
// a Signal that scales the x of another Signal
type Squeeze struct {
	LimitedSignal
	Factor float32
}

func (s Squeeze) Call(t x) y {
	return s.LimitedSignal.Call(x(float32(t) * s.Factor))
}

func (s Squeeze) MaxX() x {
	return x(float32(s.LimitedSignal.MaxX())*s.Factor)
}

*/

type SpedupProgressive struct {
	Signal
	Rate x
}

func (s SpedupProgressive) property(t x) y {
	return s.Signal.property(t + t*t/s.Rate)
}

// a PeriodicSignal that is a Signal repeated with Loop length x.
type Looped struct {
	Signal
	Loop x
}

func (s Looped) property(t x) y {
	return s.Signal.property(t % s.Loop)
}

func (s Looped) Period() x {
	return s.Loop
}

// a PeriodicSignal that is repeating loop of Cycle repeats of another PeriodicSignal.
// if the PeriodicSignal is actually precisely repeating, then an integer value of Cycles, results in no change.
type Repeated struct {
	PeriodicSignal
	Cycles float32
}

func (s Repeated) Period() x {
	return x(float32(s.PeriodicSignal.Period()) * s.Cycles)
}

func (s Repeated) property(t x) y {
	return s.PeriodicSignal.property((t % s.Period()) % s.PeriodicSignal.Period())
}

// a Signal that produces y values that are the negative of another Signals y values
type Inverted struct {
	Signal
}

func (s Inverted) property(t x) y {
	return -s.Signal.property(t)
}

// a Signal that returns y's that are for the -ve x of another Signal
type Reversed struct {
	Signal
}

func (s Reversed) property(t x) y {
	return s.Signal.property(-t)
}

// a Signal that produces values that are flipped over, (Maxy<->zero) of another Signal
type Reflected struct {
	Signal
}

func (s Reflected) property(t x) y {
	if r := s.Signal.property(t); r < 0 {
		return -unitY - r
	} else {
		return unitY - r
	}
}

// a Signal that stretches the x values of another Signal, in proportion to the value of a modulation Signal
type RateModulated struct {
	Signal
	Modulation Signal
	Factor     x
}

func (s RateModulated) property(t x) y {
	return s.Signal.property(t + MultiplyX(float64(s.Modulation.property(t))/maxyfloat64, s.Factor))
}

// Segmented is a Signal that is a sequence of equal width, uniform gradient, segments, that approximate another Signal.
// repeated calls within the same segment, are generated from cached end values, so avoiding calls to the embedded Signal.
type Segmented struct {
	Signal
	Width x
	cache *segmentCache // by being a pointer this is mutable in methods, without needing a pointer receiver.

}

type segmentCache struct {
	x1, x2 x
	l1, l2 x
}

func NewSegmented(s Signal, w x) Segmented {
	return Segmented{s, w, &segmentCache{}}
}

func (s Segmented) property(t x) y {
	temp := t % s.Width
	if t-temp != s.cache.x1 || t+s.Width-temp != s.cache.x2 {
		// TODO reuse by swap ends
		s.cache.x1 = t - temp
		s.cache.x2 = t + s.Width - temp
		s.cache.l1 = x(s.Signal.property(s.cache.x1)) / s.Width
		s.cache.l2 = x(s.Signal.property(s.cache.x2))/s.Width - s.cache.l1
	}
	return y(s.cache.l1*s.Width + s.cache.l2*temp)
}

// Triggered shifts a Signal's x so the Signal crosses a trigger y at zero x.
// it searches with a Resolution, from Shift+Resolution to MaxShift, then from 0 to Shift.
// Shift can be set initially, then is set to the last found trigger, so subsequent uses find new crossings, and wraps round.
// Rising can be alternated to find either way crossing
type Triggered struct {
	Signal
	Trigger    y
	Rising     bool
	Resolution x
	MaxShift   x
	Found      *searchInfo // by being a pointer this is mutable in methods, without needing a pointer receiver.
}

type searchInfo struct {
	Shift   x
	trigger y
	rising  bool
}

func NewTriggered(s Signal, trigger y, rising bool, res, max x) Triggered {
	return Triggered{s, trigger, rising, res, max, &searchInfo{}}
}

func (s Triggered) property(t x) y {
	if s.Trigger != s.Found.trigger || s.Found.rising != s.Rising {
		s.Found.trigger = s.Trigger
		s.Found.rising = s.Rising
		if s.Rising && s.Signal.property(s.Found.Shift) > s.Trigger || !s.Rising && s.Signal.property(s.Found.Shift) < s.Trigger {
			s.Found.Shift += s.Resolution
		}
		for t := s.Found.Shift; t <= s.MaxShift; t += s.Resolution {
			if s.Rising && s.Signal.property(t) > s.Trigger || !s.Rising && s.Signal.property(t) < s.Trigger {
				s.Found.Shift = t
				return s.Signal.property(t)
			}
		}
		for t := x(0); t < s.Found.Shift; t += s.Resolution {
			if s.Rising && s.Signal.property(t) > s.Trigger || !s.Rising && s.Signal.property(t) < s.Trigger {
				s.Found.Shift = t
				return s.Signal.property(t)
			}
		}
		s.Found.Shift = 0
	}
	return s.Signal.property(t + s.Found.Shift)
}
