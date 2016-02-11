package signals

func MultiplyInterval(m interface{}, d Interval) Interval {
	switch mt := m.(type) {
	case int:
		return d * Interval(mt)
	case uint:
		return d * Interval(mt)
	case int8:
		return d * Interval(mt)
	case uint8:
		return d * Interval(mt)
	case int16:
		return d * Interval(mt)
	case uint16:
		return d * Interval(mt)
	case int32:
		return d * Interval(mt)
	case uint32:
		return d * Interval(mt)
	case int64:
		return d * Interval(mt)
	case uint64:
		return d * Interval(mt)
	case float32:
		return Interval(float32(d)*mt + .5)
	case float64:
		return Interval(float64(d)*mt + .5)
	default:
		return d
	}
}

