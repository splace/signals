package signals

import (
	"fmt"
	"math"
	"time"
)

// function types can represent an analogue y as it varies with x
type Function interface {
	property(x) y
}

// returns a PeriodicLimitedFunction (type Modulated) based on a sine wave,
// with peak y set to Maxy adjusted by dB,
// so dB should always be negative.
func NewTone(period x, dB float64) Modulated {
	return Modulated{Sine{period}, NewConstant(dB)}
}

// the x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
// -ve x's are considered imaginary, not used, unless a Delay makes them +ve.
type x time.Duration // current underlying representation

func (i x) String() string {
	return fmt.Sprintf("%9.2f", float32(i)/float32(unitX))
}

// somewhere close to the middle of the resolution range.
const unitX = x(time.Second)

/*
// the y type represents a value between +maxY and -maxY.
type y int32

const maxY y = math.MaxInt32
const yBits = 32
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const maxyfloat64 float64 = float64(maxY - 64)   // 512
*/

// the y type represents a value between +maxY and -maxY.
type y int64

const unitY y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const maxyfloat64 float64 = float64(unitY - 512)

func (l y) String() string {
	return fmt.Sprintf("%7.2f%%", 100*float32(l)/float32(unitY))
}

// LimitedFunctions are used as Functions that can be assumed is zero after MaxX
type LimitedFunction interface {
	Function
	MaxX() x
}

// Samples are LimitedFunctions that are assumed to be zero before their MinX.
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

// Converters to promote slices of interfaces, needed when using variadic parameters called using a slice since go doesn't automatically promote a narrow interface inside the slice to be able to use a broader interface.
// for example: without these you couldn't use a slice of LimitedFunction's in a variadic call to a func requiring Function's. (when you can use separate LimitedFunction's in the same call.)

// converts  to []Function
func PromoteToFunctions(s interface{}) (out []Function) {
	switch st := s.(type) {
	case []LimitedFunction:
		out = make([]Function, len(st))
		for i := range out {
			out[i] = st[i].(Function)
		}
	case []PeriodicLimitedFunction:
		out = make([]Function, len(st))
		for i := range out {
			out[i] = st[i].(Function)
		}
	case []PeriodicFunction:
		out = make([]Function, len(st))
		for i := range out {
			out[i] = st[i].(Function)
		}
	case []PCMFunction:
		out = make([]Function, len(st))
		for i := range out {
			out[i] = st[i].(Function)
		}
	}
	return
}

// converts to []LimitedFunction
func PromoteToLimitedFunctions(s interface{}) (out []LimitedFunction) {
	switch st := s.(type) {
	case []PeriodicLimitedFunction:
		out = make([]LimitedFunction, len(st))
		for i := range out {
			out[i] = st[i].(LimitedFunction)
		}
	case []PCMFunction:
		out = make([]LimitedFunction, len(st))
		for i := range out {
			out[i] = st[i].(LimitedFunction)
		}
	}
	return
}

// converts to []PeriodicFunction
func PromoteToPeriodicFunctions(s interface{}) (out []PeriodicFunction) {
	switch st := s.(type) {
	case []PeriodicLimitedFunction:
		out = make([]PeriodicFunction, len(st))
		for i := range out {
			out[i] = st[i].(PeriodicFunction)
		}
	case []PCMFunction:
		out = make([]PeriodicFunction, len(st))
		for i := range out {
			out[i] = st[i].(PeriodicFunction)
		}
	}
	return
}

// converts to []PeriodicLimitedFunction
func PromoteToPeriodicLimitedFunctions(s interface{}) (out []PeriodicLimitedFunction) {
	switch st := s.(type) {
	case []PCMFunction:
		out = make([]PeriodicLimitedFunction, len(st))
		for i := range out {
			out[i] = st[i].(PeriodicLimitedFunction)
		}
	}
	return
}

