package signals

import (
	"fmt"
	"math"
	"time"
)

// function types can represent an analogue y as it varies with x
type Function interface {
	call(x) y
}

// x is from -infinity to +infinity, can be considered a time Dx.
// y's at -v'e xs are considered kind of imaginary, and not used, unless a Delay makes them +ve.
type x time.Duration

// formatted representation of an x.
func (i x) String() string {
	return fmt.Sprintf("%9.2f", float32(i)/float32(unitX))
}

// use time.Second as sensible middle of resolution range of int64.
const unitX = x(time.Second)

// the y type is a value between +Maxy and -Maxy.
type y int64

const maxY y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const Maxyfloat64 float64 = float64(maxY - 512)

// formatted representation of a y as percentage.
func (l y) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(maxY))
}

// Periodic's have a Period method requirment.
type Periodic interface {
	Period() x
}

// peaker adds a PeakY method requirement.
type peaker interface {
	PeakY() y
}

// LimitedFunctions are used as Functions that can be assumed is zero after MaxX
type LimitedFunction interface {
	Function
	MaxX() x
}

// Samples are LimitedFunctions that can be assumed to be zero before their MinX.
type Sample interface {
	LimitedFunction
	MinX() x
}

// PeriodicalFunctions are functions that can be assumed to repeat over Period().
type PeriodicFunction interface {
	Function
	Periodic
}

// LimitedPeriodicalFunction are Functions that repeat over Period() and dont exceed MaxY().
type PeriodicLimitedFunction interface {
	LimitedFunction
	Periodic
}

// PeakingLimitedFunctions are Functions that can be assumed is zero after a MaxX() and dont exceed a MaxY().
type PeakingLimitedFunction interface {
	LimitedFunction
	peaker
}

// PeakingLimitedPeriodicalFunction are Functions that can be assumed is zero after a MaxX(), repeat over Period() and dont exceed MaxY().
type PeakingPeriodicLimitedFunction interface {
	LimitedFunction
	Periodic
	peaker
}

// Converters to promote slices of interfaces, needed when using variadic parameters called using a slice since go doesn't automatically promote a narrow interface inside the slice to be able to use a broader interface.
// for example: without these you couldn't use a slice of LimitedFunction's in a variadic call to a func requiring Function's. (when you can use separate LimitedFunction's in the same call.)

// converts []LimitedFunction to []Function
func LimitedFunctionsToSliceFunction(s ...LimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []Function
func PeriodicLimitedFunctionsToSliceFunction(s ...PeriodicLimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicFunction to []Function
func PeriodicFunctionsToSliceFunction(s ...PeriodicFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []LimitedFunction
func PeriodicLimitedFunctionsToSliceLimitedFunction(s ...PeriodicLimitedFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []Function
func PCMFunctionsToSliceFunction(s ...PCMFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PCMFunction to []LimitedFunction
func PCMFunctionsToSliceLimitedFunction(s ...PCMFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicLimitedFunction
func PCMFunctionsToSlicePeriodicLimitedFunction(s ...PCMFunction) []PeriodicLimitedFunction {
	out := make([]PeriodicLimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicLimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicFunction
func PCMFunctionsToSlicePeriodicFunction(s ...PCMFunction) []PeriodicFunction {
	out := make([]PeriodicFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicFunction)
	}
	return out
}
