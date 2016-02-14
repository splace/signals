package signals

import (
	"fmt"
	"math"
	"time"
)

// signal types can represent an analogue level as it varies with time
type Signal interface {
	Level(interval) level
}

// Interval is considered to be a time duration from -infinity to +infinity.
// Intervals here can be generated from time.Duration, signals.Interval(time.Duration).
// encoded as a time.Duration, which is encoded as an int64, giving actually a range of 290 years at nanosecond resolution.
// Levels at -ve intervals are considered imaginary, and not used, unless a Delay makes them +ve.
type interval time.Duration

func (i interval) String() string {
	return fmt.Sprintf("%9.2fs", float32(i)/float32(UnitTime))
}

const UnitTime = interval(time.Second)

// the Level type is a value between +MaxLevel and -MaxLevel.
type level int64

const MaxLevel level = math.MaxInt64
const LevelBits = 64
const HalfLevelBits = LevelBits / 2

//const HalfLevel=2<<(HalfLevelBits-1)

// float64 has less resolution than int64 at maxlevel, so need this to scale float64 sourced signals to never overflow int64
const MaxLevelfloat64 float64 = float64(MaxLevel - 512)

// formatted representation of a level as percentage.
func (l level) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(MaxLevel))
}

