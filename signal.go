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

// returns a PeriodicLimitedFunction (type multiplex) based on a sine wave,
// with peak y set to Maxy adjusted by dB,
// so dB should always be negative.
func NewTone(period x, dB float64) Multiplex {
	return Multiplex{Sine{period}, NewConstant(dB)}
}

// x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
// -ve x's are considered imaginary, not used, unless a Delay makes them +ve.
type x time.Duration // current underlying representation

func (i x) String() string {
	return fmt.Sprintf("%9.2f", float32(i)/float32(unitX))
}

// use time.Second as sensible middle of resolution range of time.Duration.
const unitX = x(time.Second)

// the y type represents a value between +maxY and -maxY.
type y int64

const maxY y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const maxyfloat64 float64 = float64(maxY - 512)

func (l y) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(maxY))
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
	Period() x
}

// LimitedPeriodicalFunction are Functions that repeat over Period() and dont exceed MaxY().
type PeriodicLimitedFunction interface {
	LimitedFunction
	Period() x
}

// PeakingLimitedFunctions are Functions that can be assumed is zero after a MaxX() and dont exceed a MaxY().
type PeakingLimitedFunction interface {
	LimitedFunction
	peaker
}

// PeakingLimitedPeriodicalFunction are Functions that can be assumed is zero after a MaxX(), repeat over Period() and dont exceed MaxY().
type PeakingPeriodicLimitedFunction interface {
	LimitedFunction
	Period() x
	peaker
}

// Converters to promote slices of interfaces, needed when using variadic parameters called using a slice since go doesn't automatically promote a narrow interface inside the slice to be able to use a broader interface.
// for example: without these you couldn't use a slice of LimitedFunction's in a variadic call to a func requiring Function's. (when you can use separate LimitedFunction's in the same call.)

// converts []LimitedFunction to []Function
func PromoteSliceLimitedFunctionsToFunctions(s []LimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []Function
func PromoteSlicePeriodicLimitedFunctionsToFunctions(s []PeriodicLimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicFunction to []Function
func PromoteSlicePeriodicFunctionsToFunctions(s []PeriodicFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []LimitedFunction
func PromoteSlicePeriodicLimitedFunctionsToLimitedFunctions(s []PeriodicLimitedFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []Function
func PromoteSlicePCMFunctionsToFunctions(s []PCMFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PCMFunction to []LimitedFunction
func PromoteSlicePCMFunctionsToLimitedFunctions(s []PCMFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicLimitedFunction
func PromoteSlicePCMFunctionsToPeriodicLimitedFunctions(s []PCMFunction) []PeriodicLimitedFunction {
	out := make([]PeriodicLimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicLimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicFunction
func PromoteSlicePCMFunctionsToPeriodicFunctions(s []PCMFunction) []PeriodicFunction {
	out := make([]PeriodicFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicFunction)
	}
	return out
}
