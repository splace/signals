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

func (s Constant) Call(t x) y {
	return s.Constant
}

func DB(vol float64) float64 {
	return 6 * math.Log2(vol)
}
func Vol(DB float64) float64 {
	return math.Pow(2, DB/6)
}

func NewConstant(DB float64) Constant {
	return Constant{y(Maxyfloat64 * Vol(DB))}
}

type Sine struct {
	Cycle x
}

func (s Sine) Call(t x) y {
	return y(math.Sin(float64(t)/float64(s.Cycle)*2*math.Pi) * Maxyfloat64)
}

func (s Sine) Period() x {
	return s.Cycle
}

type Pulse struct {
	Width x
}

func (s Pulse) Call(t x) y {
	if t > s.Width || t < 0 {
		return 0
	} else {
		return maxy
	}
}

func (s Pulse) MaxX() x {
	return s.Width
}

type Square struct {
	Cycle x
}

func (s Square) Call(t x) y {
	if t%s.Cycle >= s.Cycle/2 {
		return -maxy
	} else {
		return maxy
	}
}

func (s Square) Period() x {
	return s.Cycle
}

type RampUp struct {
	Period x
}

func (s RampUp) Call(t x) y {
	if t < 0 {
		return 0
	} else if t > s.Period {
		return maxy
	} else {
		return y(x(maxy) / s.Period * t)
	}
}

type RampDown struct {
	Period x
}

func (s RampDown) Call(t x) y {
	if t < 0 {
		return maxy
	} else if t > s.Period {
		return 0
	} else {
		return y(x(maxy) / s.Period * (s.Period - t))
	}
}

type Heavyside struct {
}

func (s Heavyside) Call(t x) y {
	if t < 0 {
		return 0
	}
	return maxy
}

type Sigmoid struct {
	Steepness x
}

func (s Sigmoid) Call(t x) y {
	return y(Maxyfloat64 / (1 + math.Exp(-float64(t)/float64(s.Steepness))))
}
