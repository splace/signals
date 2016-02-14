package signals

// periodicals are signals that repeat
type Periodical interface {
	Signal
	Period() interval
}

func NewTone(period interval, volume uint8) Multi {
	return Multi{Sine{period}, NewConstant(volume)}
}

/*
type Tone struct{
	Signal
	Cycle Interval
}

func (s Tone) Period() Interval{
	return s.Cycle
}
*/
// make a periodical whose source is a sine
