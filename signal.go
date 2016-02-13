package signals

import (
	"fmt"
	"math"
	"time"
)

// signal types can represent an analogue level as it varies with time
type Signal interface {
	Level(Interval) Level
}

// Interval is considered to be a time duration from -infinity to +infinity.
// Intervals here can be generated from time.Duration, signals.Interval(time.Duration).
// encoded as a time.Duration, which is encoded as an int64, giving actually a range of 290 years at nanosecond resolution.
// Levels at -ve intervals are considered imaginary, and not used, unless a Delay makes them +ve.
type Interval time.Duration

func (i Interval) String() string {
	return fmt.Sprintf("%9.2fs", float32(i)/float32(UnitTime))
}

const UnitTime = Interval(time.Second)

// the Level type is a value between +MaxLevel and -MaxLevel.
type Level int64

const MaxLevel Level = math.MaxInt64
const LevelBits = 64
const HalfLevelBits = LevelBits / 2

//const HalfLevel=2<<(HalfLevelBits-1)

// float64 has less resolution than int64 at maxlevel, so need this to scale float64 sourced signals to never overflow int64
const MaxLevelfloat64 float64 = float64(MaxLevel - 512)

// formatted representation of a level as percentage.
func (l Level) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(MaxLevel))
}


