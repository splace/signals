package signals

import "time"

func X(d time.Duration) x {
	return MultiplyX(d.Seconds(), UnitX)
}

func MultiplyX(m interface{}, d x) x {
	switch mt := m.(type) {
	case int:
		return d * x(mt)
	case uint:
		return d * x(mt)
	case int8:
		return d * x(mt)
	case uint8:
		return d * x(mt)
	case int16:
		return d * x(mt)
	case uint16:
		return d * x(mt)
	case int32:
		return d * x(mt)
	case uint32:
		return d * x(mt)
	case int64:
		return d * x(mt)
	case uint64:
		return d * x(mt)
	case float32:
		return x(float32(d)*mt + .5)
	case float64:
		return x(float64(d)*mt + .5)
	default:
		return d
	}
}
