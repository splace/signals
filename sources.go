package signals

import "encoding/gob"

import (
	"math"
)

func init() {
	gob.Register(Constant{})
	gob.Register(Sine{})
	gob.Register(Sinc{})
	gob.Register(Pulse{})
	gob.Register(Square{})
	gob.Register(RampUp{})
	gob.Register(RampDown{})
	gob.Register(Heavyside{})
	gob.Register(Sigmoid{})
	gob.Register(Gauss{})
}

func DB(vol float64) float64 {
	return 6 * math.Log2(vol)
}
func Vol(DB float64) float64 {
	return math.Pow(2, DB/6)
}

// a Signal with constant value
type Constant struct {
	Constant y
}

func NewConstant(DB float64) Constant {
	return Constant{y(unitYfloat64 * Vol(DB))}
}

func (s Constant) property(p x) y {
	return s.Constant
}

// a PeriodicSignal that varies sinusoidally, repeating with Cycle width.
type Sine struct {
	Cycle x
}

func (s Sine) property(p x) y {
	return y(math.Sin(float64(p)/float64(s.Cycle)*2*math.Pi) * unitYfloat64)
}

func (s Sine) Period() x {
	return s.Cycle
}

// a Signal that is a wave 'packet', with fundamental (central) wavelength of Cycle.
type Sinc struct {
	Cycle x
}

func (s Sinc) property(p x) y {
	if p==0 {return unitY}
	xp:=float64(p)/float64(s.Cycle)*2*math.Pi
	return y(math.Sin(xp)/xp * unitYfloat64)
}

// a Signal that peaks, centred on zero, as a Gaussian distribution, width q.
type Gauss struct {
	Q22 float64   // 2 * q squared
}

func (s Gauss) property(p x) y {
	return y(math.Exp(-float64(p)*float64(p)/s.Q22) * unitYfloat64)
}


// a LimitedSignal that produces unitY for a Width, zero otherwise.
type Pulse struct {
	Width x
}

func (s Pulse) property(p x) y {
	if p > s.Width || p < 0 {
		return 0
	} else {
		return unitY
	}
}

func (s Pulse) MaxX() x {
	return s.Width
}

// a PeriodicSignal that produces equal regions of +unitY then -unitY, repeating with Cycle width.
type Square struct {
	Cycle x
}

func (s Square) property(p x) y {
	if p%s.Cycle >= s.Cycle/2 {
		return -unitY
	} else {
		return unitY
	}
}

func (s Square) Period() x {
	return s.Cycle
}

// a Signal which ramps from zero to unitY over a Period width.
type RampUp struct {
	Period x
}

func (s RampUp) property(p x) y {
	if p < 0 {
		return 0
	} else if p > s.Period {
		return unitY
	} else {
		return y(x(unitY) / s.Period * p)
	}
}

// a Signal which ramps from unitY to zero, over a Period width.
type RampDown struct {
	Period x
}

func (s RampDown) property(p x) y {
	if p < 0 {
		return unitY
	} else if p > s.Period {
		return 0
	} else {
		return y(x(unitY) / s.Period * (s.Period - p))
	}
}

// a Signal that returns +unitY for positive x and zero for negative x.
type Heavyside struct {
}

func (s Heavyside) property(p x) y {
	if p < 0 {
		return 0
	}
	return unitY
}

// a Signal that smoothly transitions from 0 to +unitY.
// with a maximum gradient (first derivative) at x=0, of Steepness.
type Sigmoid struct {
	Steepness x
}

func (s Sigmoid) property(p x) y {
	return y(unitYfloat64 / (1 + math.Exp(-float64(p)/float64(s.Steepness))))
}


