package signals
import	"encoding/gob"

import (
	"math"
)
func init() {
	gob.Register(Sine{})
	gob.Register(Constant{})
	gob.Register(Pulse{})
	gob.Register(Square{})
	gob.Register(RampUp{})
	gob.Register(RampDown{})
	gob.Register(Heavyside{})
	gob.Register(Sigmoid{})
}

type Constant struct {
	Constant level "level"
}

func (s Constant) Level(t interval) level {
	return s.Constant
}

func NewConstant(volume uint8) Constant {
	return Constant{MaxLevel / 100 * level(volume)}
}

type Sine struct {
	Cycle interval
}

func (s Sine) Level(t interval) level {
	return level(math.Sin(float64(t)/float64(s.Cycle)*2*math.Pi) * MaxLevelfloat64)
}

func (s Sine) Period() interval {
	return s.Cycle
}

type Pulse struct {
	Width interval
}

func (s Pulse) Level(t interval) level {
	if t > s.Width {
		return 0
	} else {
		return MaxLevel
	}
}

type Square struct {
	Cycle interval
}

func (s Square) Level(t interval) level {
	if t%s.Cycle >= s.Cycle/2 {
		return -MaxLevel
	} else {
		return MaxLevel
	}
}

func (s Square) Period() interval {
	return s.Cycle
}

type RampUp struct {
	Period interval
}

func (s RampUp) Level(t interval) level {
	if t < 0 {
		return 0
	} else if t > s.Period {
		return MaxLevel
	} else {
		return level(interval(MaxLevel) / s.Period * t)
	}
}

type RampDown struct {
	Period interval
}

func (s RampDown) Level(t interval) level {
	if t < 0 {
		return MaxLevel
	} else if t > s.Period {
		return 0
	} else {
		return level(interval(MaxLevel) / s.Period * (s.Period - t))
	}
}

type Heavyside struct {
}

func (s Heavyside) Level(t interval) level {
	if t < 0 {
		return 0
	}
	return MaxLevel
}

type Sigmoid struct {
	Steepness interval
}

func (s Sigmoid) Level(t interval) level {
	return level(float64(MaxLevel) / (1 + math.Exp(-float64(t)/float64(s.Steepness))))
}
