package signals

import (
	"fmt"
	"math"
	"time"
)

// TODO cache

// signal types can represent an analogue level as it varies with time
type Signal interface {
	Level(interval) level
}

// Interval is from -infinity to +infinity, consider it a time duration.
// Levels at -ve intervals are considered kind of imaginary, and not used, unless a Delay makes them +ve.
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

// limiter adds a Duration method requirment.
type limiter interface {
	Duration() interval
}

// LimitedSignals are used as Signals that can be assumed is zero after a duration
type LimitedSignal interface {
	Signal
	limiter
}

// Samples are used as LimitedSignals that can be assumed to be zero before their start interval.
type Sample interface {
	LimitedSignal
	Start() interval
}

// Periodical adds a Period method requirment.
type Periodic interface {
	Period() interval
}

// PeriodicalSignals are signals that can be assumed to repeat
type PeriodicSignal interface {
	Signal
	Periodic
}

// LimitedPeriodicalSignal are Signals that repeat and coe to an end.
type LimitedPeriodicSignal interface {
	Signal
	limiter
	Periodic
}
