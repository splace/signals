package signals

import "encoding/gob"

func init() {
	gob.Register(Shifted{})
	gob.Register(Compressed{})
	gob.Register(Looped{})
	gob.Register(Inverted{})
	gob.Register(Reversed{})
	gob.Register(Reflected{})
	gob.Register(RateModulated{})
	gob.Register(Triggered{})
	gob.Register(&Segmented{})
}

// a Signal that is the another Signal shifted
type Shifted struct {
	Signal
	Shift x
}

func (s Shifted) property(offset x) y {
	return s.Signal.property(offset - s.Shift)
}

func (s Shifted) MaxX() x {
	return s.Signal.(LimitedSignal).MaxX()+s.Shift
}

// a Signal that scales the x of another Signal
type Compressed struct {
	 Signal
	Factor float32
}

func (s Compressed) property(offset x) y {
	return s.Signal.property(x(float32(offset) * s.Factor))
}

func (s Compressed) MaxX() x {
	return x(float32(s.Signal.(LimitedSignal).MaxX())/s.Factor)
}

func (s Compressed) Period() x {
	return x(float32(s.Signal.(PeriodicSignal).Period())/s.Factor)
}

// a PeriodicSignal that is a Signal repeated with Loop length x.
type Looped struct {
	Signal
	Loop x
}

func (s Looped) property(offset x) y {
	return s.Signal.property(offset % s.Loop)
}

func (s Looped) Period() x {
	return s.Loop
}

// a PeriodicSignal that is repeating loop of Cycles number of repeats of another PeriodicSignal.
// if the PeriodicSignal is actually precisely repeating, then an integer value of Cycles, results in no change.
type Repeated struct {
	PeriodicSignal
	Cycles float32
}

func (s Repeated) Period() x {
	return x(float32(s.PeriodicSignal.Period()) * s.Cycles)
}

func (s Repeated) property(offset x) y {
	return s.PeriodicSignal.property((offset % s.Period()) % s.PeriodicSignal.Period())
}


// a Signal that produces y values that are the negative of another Signals y values
type Inverted struct {
	Signal
}

func (s Inverted) property(offset x) y {
	return -s.Signal.property(offset)
}

// a Signal that returns y's that are for the -ve x of another Signal
type Reversed struct {
	Signal
}

func (s Reversed) property(offset x) y {
	return s.Signal.property(-offset)
}

// a Signal that produces values that are flipped over, (Maxy<->zero) of another Signal
type Reflected struct {
	Signal
}

func (s Reflected) property(offset x) y {
	if r := s.Signal.property(offset); r < 0 {
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

func (s RateModulated) property(offset x) y {
	return s.Signal.property(offset + MultiplyX(float64(s.Modulation.property(offset))/maxyfloat64, s.Factor))
}

func (s RateModulated) Period() x {
	if  mps,ok:= s.Modulation.(PeriodicSignal);ok {
		if  ps,ok:= s.Signal.(PeriodicSignal);ok {
			if mpsp,psp:=mps.Period(),ps.Period(); mpsp>psp {
				return mpsp
			}else{ 
				return psp
			}	
		}else{
			return mps.Period()
		}
	}else{
		s.Signal.(PeriodicSignal).Period()
	}
	return 0
}


// Segmented is a Signal that is a sequence of equal width, uniform gradient, segments, that approximate another Signal.
// repeated calls within the same segment, are generated from cached end values, so avoiding calls to the embedded Signal.
type Segmented struct {
	Signal
	Width x
	x1, x2, l1, l2 x
}


func (s *Segmented) property(offset x) y {
	temp := offset % s.Width
	if offset-temp != s.x1 || offset+s.Width-temp != s.x2 {
		// TODO reuse by swap ends
		s.x1 = offset - temp
		s.x2 = offset + s.Width - temp
		s.l1 = x(s.Signal.property(s.x1)) / s.Width
		s.l2 = x(s.Signal.property(s.x2))/s.Width - s.l1
	}
	return y(s.l1*s.Width + s.l2*temp)
}

func (s Segmented) Period() x {
	// TODO could use shortest of width and signal periods?
	return s.Width
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

func (s Triggered) property(offset x) y {
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
	return s.Signal.property(offset + s.Found.Shift)
}



