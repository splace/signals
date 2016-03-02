package signals

import "time"

// allows using x without direct access 
func X(d time.Duration) x {
	return MultiplyX(d.Seconds(), UnitX)
}

// allows using x without direct access 
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

// series of converts to promote slicea of interfaces. 

// converts []LimitedFunction to []Function 
func LimitedFunctionsToSliceFunction(s ...LimitedFunction) []Function{
	out:=make([]Function,len(s))
	for i:=range(out){
		out[i]=s[i].(Function)
	}
	return out
} 
// converts []LimitedFunction to []Function 
func PeriodicLimitedFunctionsToSliceFunction(s ...PeriodicLimitedFunction) []Function{
	out:=make([]Function,len(s))
	for i:=range(out){
		out[i]=s[i].(Function)
	}
	return out
} 


// converts []LimitedFunction to []Function 
func PeriodicLimitedFunctionsToSliceLimitedFunction(s ...PeriodicLimitedFunction) []LimitedFunction{
	out:=make([]LimitedFunction,len(s))
	for i:=range(out){
		out[i]=s[i].(LimitedFunction)
	}
	return out
} 

// converts []LimitedFunction to []Function 
func PCMFunctionsToSliceFunction(s ...PCMFunction) []Function{
	out:=make([]Function,len(s))
	for i:=range(out){
		out[i]=s[i].(Function)
	}
	return out
} 
// converts []LimitedFunction to []Function 
func PCMFunctionsToSliceLimitedFunction(s ...PCMFunction) []LimitedFunction{
	out:=make([]LimitedFunction,len(s))
	for i:=range(out){
		out[i]=s[i].(LimitedFunction)
	}
	return out
} 

// converts []LimitedFunction to []Function 
func PCMFunctionsToSlicePeriodicLimitedFunction(s ...PCMFunction) []PeriodicLimitedFunction{
	out:=make([]PeriodicLimitedFunction,len(s))
	for i:=range(out){
		out[i]=s[i].(PeriodicLimitedFunction)
	}
	return out
} 



