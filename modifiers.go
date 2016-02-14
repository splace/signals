package signals

// a Signal that delays the time of another signal
type Delayed struct {
	Signal
	Delay interval
}

func (s Delayed) Level(t interval) level {
	return s.Signal.Level(t - s.Delay)
}

// a Signal that sppeds up the time of another signal
type Spedup struct {
	Signal
	Factor float32
}

func (s Spedup) Level(t interval) level {
	return s.Signal.Level(interval(float32(t) * s.Factor))
}

// TODO spedup tone should have period changed

type SpedupProgressive struct {
	Signal
	Rate interval
}

func (s SpedupProgressive) Level(t interval) level {
	return s.Signal.Level(t + t*t/s.Rate)
}

// a Signal that repeats another signal
type Looped struct {
	Signal
	Length interval
}

func (s Looped) Level(t interval) level {
	return s.Signal.Level(t % s.Length)
}

func (s Looped) Period() interval {
	return s.Length
}

// a Signal that produced level values that are the negative of another signals level values
type Inverted struct {
	Signal
}

func (s Inverted) Level(t interval) level {
	return -s.Signal.Level(t)
}

// a Signal that returns levels that run time backwards of another signal
type Reversed struct {
	Signal
}

func (s Reversed) level(t interval) level {
	return s.Signal.Level(-t)
}

// a Signal that produces values that are flipped over, (Maxvalue<->zero) of another signal
type Reflected struct {
	Signal
}

func (s Reflected) Level(t interval) level {
	if r := s.Signal.Level(t); r < 0 {
		return -MaxLevel - r
	} else {
		return MaxLevel - r
	}
}

// a Signal that stretches the time of another signal, in proportion to the value of a modulation signal
type Modulated struct {
	Signal
	Modulation Signal
	Factor     interval
}

func (s Modulated) Level(t interval) level {
	return s.Signal.Level(t + MultiplyInterval(float64(s.Modulation.Level(t))/MaxLevelfloat64, s.Factor))
}

// brings forward in time a signal to the point when it passes a trigger level at zero time.
type Triggered struct {
	Signal
	Trigger        level
	Rising	bool
	Resolution     interval
	MaxDelay       interval
	Delay          interval
	searched       Signal
	locatedTrigger level
}

func (s Triggered) Level(t interval) level {
	if s.Trigger != s.locatedTrigger || s.searched != s.Signal {
		s.searched = s.Signal
		s.locatedTrigger = s.Trigger
		s.Delay = 0
		for t := interval(0); t < s.MaxDelay; t += s.Resolution {
			if s.Rising && s.Signal.Level(t) > s.Trigger || !s.Rising && s.Signal.Level(t) < s.Trigger {
				s.Delay = t
				break
			}
		}
	}
	return s.Signal.Level(t + s.Delay)
}
