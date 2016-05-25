package signals

import (
	"fmt"
	"math"
	"time"
)

// satisfying the Function interface means a type represents an analogue signal, where a y property varies with an x parameter.
type Signal interface {
	property(x) y
}

// the x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
// -ve x's are considered imaginary, not used, unless a Delay makes them +ve.
type x time.Duration // current underlying representation

// somewhere close to the middle of the resolution range.
const unitX = x(time.Second)

// string representation of an x scaled to unitX
func (i x) String() string {
	return fmt.Sprintf("%9.2f", float32(i)/float32(unitX))
}

// the y type represents a value between +unitY and -unitY.
type y int64

const unitY y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const maxyfloat64 float64 = float64(unitY - 512)

// string representation of a y, scaled to unitY%
func (l y) String() string {
	return fmt.Sprintf("%7.2f%%", 100*float32(l)/float32(unitY))
}

// LimitedFunctions are Functions that are assumed to have zero y after MaxX
type LimitedSignal interface {
	Signal
	MaxX() x
}

/*
// Samples are LimitedFunctions that are assumed to be zero before their MinX.
type Sample interface {
	LimitedFunction
	MinX() x
}*/

// PeriodicalFunctions are functions that repeat, give the same y, if x changes by the amount returned by Period().
type PeriodicSignal interface {
	Signal
	Period() x
}

// LimitedPeriodicalFunction are Functions that repeat over Period() and dont exceed MaxY().
type PeriodicLimitedSignal interface {
	LimitedSignal
	Period() x
}
