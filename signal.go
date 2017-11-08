package signals

import (
	"fmt"
	"math"
)

// satisfying the Signal interface means a type represents an analogue signal, where property of type y, varies with a parameter of type x.
type Signal interface {
	property(x) y
}

// the x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
// -ve x's are considered imaginary, not used, unless a Delay makes them +ve.
type x int64 // current underlying representation
const xBits=64
// somewhere close to the middle of the resolution range, adjust if dealing with high accuracy but only either small or large values.
const unitX = x(1000000000)

// string representation of an x scaled to unitX
func (p x) String() string {
	return fmt.Sprintf("%9.2f", float32(p)/float32(unitX))
}

// the y type represents a value between +unitY and -unitY.
type y int64

const unitY y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced Signals to never overflow int64
const unitYfloat64 float64 = float64(unitY - 512)

// string representation of a y, scaled to unitY%
func (v y) String() string {
	return fmt.Sprintf("%7.2f%%", 100*float32(v)/float32(unitY))
}

// a LimitedSignal is a Signal modified to give property values of zero with parameter values above the values returned by MaxX().
type LimitedSignal interface {
	Signal
	MaxX() x
}

// a PeriodicSignal is a Signal that repeats, that is, gives the same value of its property, for parameter values offset by the value returned by Period().
type PeriodicSignal interface {
	Signal
	Period() x
}

// a PeriodicLimitedSignal is a Signal that repeats over Period() and is zero above MaxX().
type PeriodicLimitedSignal interface {
	Signal
	MaxX() x
	Period() x
}
