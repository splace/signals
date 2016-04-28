package signals

import "encoding/gob"

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
	Constant y
}

func (s Constant) property(t x) y {
	return s.Constant
}

func DB(vol float64) float64 {
	return 6 * math.Log2(vol)
}
func Vol(DB float64) float64 {
	return math.Pow(2, DB/6)
}

func NewConstant(DB float64) Constant {
	return Constant{y(maxyfloat64 * Vol(DB))}
}

type Sine struct {
	Cycle x
}

func (s Sine) property(t x) y {
	return y(math.Sin(float64(t)/float64(s.Cycle)*2*math.Pi) * maxyfloat64)
}

func (s Sine) Period() x {
	return s.Cycle
}

type Pulse struct {
	Width x
}

func (s Pulse) property(t x) y {
	if t > s.Width || t < 0 {
		return 0
	} else {
		return unitY
	}
}

func (s Pulse) MaxX() x {
	return s.Width
}

type Square struct {
	Cycle x
}

func (s Square) property(t x) y {
	if t%s.Cycle >= s.Cycle/2 {
		return -unitY
	} else {
		return unitY
	}
}

func (s Square) Period() x {
	return s.Cycle
}

type RampUp struct {
	Period x
}

func (s RampUp) property(t x) y {
	if t < 0 {
		return 0
	} else if t > s.Period {
		return unitY
	} else {
		return y(x(unitY) / s.Period * t)
	}
}

type RampDown struct {
	Period x
}

func (s RampDown) property(t x) y {
	if t < 0 {
		return unitY
	} else if t > s.Period {
		return 0
	} else {
		return y(x(unitY) / s.Period * (s.Period - t))
	}
}

type Heavyside struct {
}

func (s Heavyside) property(t x) y {
	if t < 0 {
		return 0
	}
	return unitY
}

type Sigmoid struct {
	Steepness x
}

func (s Sigmoid) property(t x) y {
	return y(maxyfloat64 / (1 + math.Exp(-float64(t)/float64(s.Steepness))))
}
