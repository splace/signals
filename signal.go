package signals

import (
	"fmt"
	"math"
	"time"
)

// TODO cache

// function types can represent an analogue y as it varies with x
type Function interface {
	Call(x) y
}

// x is from -infinity to +infinity, can be considered a time Dx.
// y's at -v'e xs are considered kind of imaginary, and not used, unless a Delay makes them +ve.
type x time.Duration

func (i x) String() string {
	return fmt.Sprintf("%9.2fs", float32(i)/float32(UnitX))
}

const UnitX = x(time.Second)

// the y type is a value between +Maxy and -Maxy.
type y int64

const Maxy y = math.MaxInt64
const yBits = 64
const HalfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const Maxyfloat64 float64 = float64(Maxy - 512)

// formatted representation of a y as percentage.
func (l y) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(Maxy))
}

// limiter adds a MaxX method requirement.
type limiter interface {
	MaxX() x
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
	limiter
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

