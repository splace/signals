package signals

// Tones are signals that repeat periodically
type Periodical interface {
	Signal
	Period() Interval
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
// make a Tone whose source is a stretched sine
func NewTone(period Interval, volume uint8) Multi {
	return Multi{Sine{period}, NewConstant(volume)}
}
