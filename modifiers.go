package signals

type Delayed struct {
	Signal
	Delay Interval
}

func (s Delayed) Level(t Interval) Level {
	return s.Signal.Level(t - s.Delay)
}

type Spedup struct {
	Signal
	Factor float32
}

func (s Spedup) Level(t Interval) Level {
	return s.Signal.Level(Interval(float32(t) * s.Factor))
}

type SpedupProgressive struct {
	Signal
	Rate Interval
}

func (s SpedupProgressive) Level(t Interval) Level {
	return s.Signal.Level(t + t*t/s.Rate)
}

type Looping struct {
	Signal
	Length Interval
}

func (s Looping) Level(t Interval) Level {
	return s.Signal.Level(t % s.Length)
}

type Inverted struct {
	Signal
}

func (s Inverted) Level(t Interval) Level {
	return -s.Signal.Level(t)
}

type Reversed struct {
	Signal
}

func (s Reversed) Level(t Interval) Level {
	return s.Signal.Level(-t)
}

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

