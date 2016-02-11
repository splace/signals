// Package signals generates and manipulates signals.
// signals here are as:- "https://en.wikipedia.org/wiki/Signal_processing"
// where they "convey information about the behavior or attributes of some phenomenon"
// here, more specifically, "any quantity exhibiting variation in time or variation in space"
// currently this package supports only 1-Dimensionsal variation.
// and for simplicity terminolology used represents analogy variation in time.
// this package is intended to be general, and so a base package for import, and used then with specific real-world quantities.
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

/* Interval is considered to be a time duration from -infinity to +infinity.
Intervals here can be generated from time.Duration, signals.Interval(time.Duration).
encoded as a time.Duration, which is encoded as an int64, giving actually a range of 290 years at nanosecond resolution.
Levels at -ve intervals are considered imaginary, and not used, unless a Delay makes them +ve.*/
type Interval time.Duration

func (i Interval) String() string {
	return fmt.Sprintf("%9.2fs", float32(i)/float32(UnitTime))
}

const UnitTime = Interval(time.Second)

// the Level type is a value between +MaxLevel to -MaxLevel.
type Level int64

const MaxLevel Level = math.MaxInt64
const LevelBits = 64
const HalfLevelBits = LevelBits / 2

//const HalfLevel=2<<(HalfLevelBits-1)
const MaxLevelfloat64 float64 = float64(MaxLevel - 512) // float64 has less resolution than int64 at maxlevel, use this to scale some signals down

// formatted representation of a level as percentage.
func (l Level) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(MaxLevel))
}

// Product is a Signal generated by multiplying together Signal(s)
// so its Level needs to run from +unity to -unity.
// so multiplication scales so that, MaxLevel*MaxLevel=MaxLevel (so MaxLevel is unity).
// this makes Product like a logic AND; all sources (at a particular momemt) need to be MaxLevel to produce a Product of MaxLevel.
// where as, ANY source at zero will generate a Product of zero.
type Product []Signal

func (c Product) Level(t Interval) (total Level) {
	total = MaxLevel
	for _, s := range c {
		l := s.Level(t)
		switch l {
		case 0:
			total = 0
			break
		case MaxLevel:
			continue
		default:
			//total = (total / HalfLevel) * (l / HalfLevel)*2
			total = (total >> HalfLevelBits) * (l >> HalfLevelBits) * 2
		}
	}
	return
}

