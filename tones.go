package signals

// Tones are signals that repeat periodically
type Tone interface{
	Signal
	Period() Interval
}

// make a Tone whose source is a scaled sine
func NewTone(period Interval, Volume uint8) Tone {
	return Product{Sine{period}, Constant{MaxLevel / 100 * Level(Volume)}}
}


