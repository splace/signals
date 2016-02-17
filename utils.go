package signals

import "time"

func Interval(d time.Duration) interval{
	return MultiplyInterval(d.Seconds(), UnitTime)
}

func MultiplyInterval(m interface{}, d interval) interval {
	switch mt := m.(type) {
	case int:
		return d * interval(mt)
	case uint:
		return d * interval(mt)
	case int8:
		return d * interval(mt)
	case uint8:
		return d * interval(mt)
	case int16:
		return d * interval(mt)
	case uint16:
		return d * interval(mt)
	case int32:
		return d * interval(mt)
	case uint32:
		return d * interval(mt)
	case int64:
		return d * interval(mt)
	case uint64:
		return d * interval(mt)
	case float32:
		return interval(float32(d)*mt + .5)
	case float64:
		return interval(float64(d)*mt + .5)
	default:
		return d
	}
}
