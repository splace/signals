package signals

// a Signal that delays the time of another signal
type Delayed struct {
	Signal
	Delay Interval
}

func (s Delayed) Level(t Interval) Level {
	return s.Signal.Level(t - s.Delay)
}

// a Signal that sppeds up the time of another signal
type Spedup struct {
	Signal
	Factor float32
}

func (s Spedup) Level(t Interval) Level {
	return s.Signal.Level(Interval(float32(t) * s.Factor))
}

// TODO spedup tone should have period changed

type SpedupProgressive struct {
	Signal
	Rate Interval
}

func (s SpedupProgressive) Level(t Interval) Level {
	return s.Signal.Level(t + t*t/s.Rate)
}

// a Signal that repeats another signal
type Looped struct {
	Signal
	Length Interval
}

func (s Looped) Level(t Interval) Level {
	return s.Signal.Level(t % s.Length)
}

func (s Looped) Period() Interval {
	return s.Length
}

// a Signal that produced level values that are the negative of another signals level values
type Inverted struct {
	Signal
}

func (s Inverted) Level(t Interval) Level {
	return -s.Signal.Level(t)
}

// a Signal that returns levels that run time backwards of another signal
type Reversed struct {
	Signal
}

func (s Reversed) Level(t Interval) Level {
	return s.Signal.Level(-t)
}

// a Signal that produces values that are flipped over, (Maxvalue<->zero) of another signal
type Reflected struct {
	Signal
}

func (s Reflected) Level(t Interval) Level {
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
	Factor     Interval
}

func (s Modulated) Level(t Interval) Level {
	return s.Signal.Level(t + MultiplyInterval(float64(s.Modulation.Level(t))/MaxLevelfloat64, s.Factor))
}

// brings forward in time a signal to the point when it passes a trigger level at zero time.
type TriggerRising struct {
	Signal
	Trigger Level
	Resolution Interval
	MaxDelay Interval
	Delay Interval
	searched Signal
	locatedTrigger Level	
}

func (s TriggerRising) Level(t Interval) Level {
	if s.Trigger!=s.locatedTrigger || s.searched!=s.Signal {
		s.searched=s.Signal
		s.locatedTrigger=s.Trigger	
		s.Delay=0
		for t:=Interval(0);t<s.MaxDelay;t+=s.Resolution{
			if s.Signal.Level(t)>s.Trigger {
				s.Delay=t
				break
			} 
		}
	}
	return s.Signal.Level(t + s.Delay)
}

/*  hal3 Sat 13 Feb 06:04:51 GMT 2016 go version go1.5.1 linux/386
FAIL	_/home/simon/Dropbox/github/working/signals [build failed]
Sat 13 Feb 06:04:51 GMT 2016 */

