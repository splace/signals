package signals

import "encoding/gob"

import (
	"math"
)

func init() {
	gob.Register(Constant{})
	gob.Register(Sine{})
	gob.Register(Tone{})
	gob.Register(Pulse{})
	gob.Register(Square{})
	gob.Register(RampUp{})
	gob.Register(RampDown{})
	gob.Register(Heavyside{})
	gob.Register(Sigmoid{})
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
	return Constant{y(maxyfloat64 * Vol(DB))}
}

func (s Constant) property(t x) y {
	return s.Constant
}

// a PeriodicSignal that varies sinusoidally, repeating with Cycle width.
type Sine struct {
	Cycle x
}

func (s Sine) property(t x) y {
	return y(math.Sin(float64(t)/float64(s.Cycle)*2*math.Pi) * maxyfloat64)
}

func (s Sine) Period() x {
	return s.Cycle
}

// a PeriodicSignal that varies sinusoidally, repeating with Cycle width, up to a max peak y.
type Tone struct {
	Cycle x
	Peak y
}

func (s Tone) property(t x) y {
	return y(float64(s.Peak)*math.Sin(float64(t)/float64(s.Cycle)*2*math.Pi) )
}

func (s Tone) Period() x {
	return s.Cycle
}

// make a Tone with peak y set to unitY (max value) adjusted by dB, so dB should not be positive.
func NewTone(period x, dB float64) Tone {
	return Tone{period, y(maxyfloat64*Vol(dB))}
}

// a LimitedSignal that produces unitY for a Width, zero otherwise.
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

// a PeriodicSignal that produces equal regions of +unitY then -unitY, repeating with Cycle width.
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

// a Signal which ramps from zero to unitY over a Period width.
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

// a Signal wcich ramps from unitY to zero, over a Period width.
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

// a Signal that returns +unitY for positive x and zero for negative.
type Heavyside struct {
}

func (s Heavyside) property(t x) y {
	if t < 0 {
		return 0
	}
	return unitY
}

// a Signal that smoothly transitions from 0 to +unitY.
// with a maximium gradient (first derivative) at x=0, of Steepness.
type Sigmoid struct {
	Steepness x
}

func (s Sigmoid) property(t x) y {
	return y(maxyfloat64 / (1 + math.Exp(-float64(t)/float64(s.Steepness))))
}


