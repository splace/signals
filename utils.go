package signals

// convert to internal y representation, 1 -> maxY
func Y(d interface{}) y {
	return MultiplyY(d, maxY)
}

// multiply anything by an y quantity
func MultiplyY(m interface{}, d y) y {
	switch mt := m.(type) {
	case int:
		return d * y(mt)
	case uint:
		return d * y(mt)
	case int8:
		return d * y(mt)
	case uint8:
		return d * y(mt)
	case int16:
		return d * y(mt)
	case uint16:
		return d * y(mt)
	case int32:
		return d * y(mt)
	case uint32:
		return d * y(mt)
	case int64:
		return d * y(mt)
	case uint64:
		return d * y(mt)
	case float32:
		return y(float32(d)*mt + .5)
	case float64:
		return y(float64(d)*mt + .5)
	default:
		return d
	}
}

// convert to internal x representation, 1 -> UnitX
func X(d interface{}) x {
	return MultiplyX(d, unitX)
}

// multiply anything by an x quantity
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

