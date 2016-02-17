package signals

// periodicals are signals that repeat
type Periodical interface {
	Signal
	Period() interval
}

// a periodical (type multi) based on a sine wave, and having a set volume%.
func NewTone(period interval, volume float32) Multiplex {
	return Multiplex{Sine{period}, NewConstant(volume)}
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
