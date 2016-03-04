package signals

import (
	"fmt"
	"math"
	"time"
)

// function types can represent an analogue y as it varies with x
type Function interface {
	Call(x) y
}

// x is from -infinity to +infinity, can be considered a time Dx.
// y's at -v'e xs are considered kind of imaginary, and not used, unless a Delay makes them +ve.
type x time.Duration

// formatted representation of an x.
func (i x) String() string {
	return fmt.Sprintf("%9.2f", float32(i)/float32(UnitX))
}

// use time.Second as sensible middle of resolution range of int64.
const UnitX = x(time.Second)

// the y type is a value between +Maxy and -Maxy.
type y int64

const maxy y = math.MaxInt64
const yBits = 64
const halfyBits = yBits / 2

//const Halfy=2<<(HalfyBits-1)

// float64 has less resolution than int64 at maxy, so need this to scale float64 sourced functions to never overflow int64
const Maxyfloat64 float64 = float64(maxy - 512)

// formatted representation of a y as percentage.
func (l y) String() string {
	return fmt.Sprintf("%9.2f%%", 100*float32(l)/float32(maxy))
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
	MaxX() x
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

// Converters to promote slices of interfaces, needed when using variadic parameters called using a slice since go doesn't automatically promote a narrow interface inside the slice to be able to use a broader interface.
// for example: without these you couldn't use a slice of LimitedFunction's in a variadic call to a func requiring Function's. (when you can use separate LimitedFunction's in the same call.)

// converts []LimitedFunction to []Function
func LimitedFunctionsToSliceFunction(s ...LimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []Function
func PeriodicLimitedFunctionsToSliceFunction(s ...PeriodicLimitedFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicFunction to []Function
func PeriodicFunctionsToSliceFunction(s ...PeriodicFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PeriodicLimitedFunction to []LimitedFunction
func PeriodicLimitedFunctionsToSliceLimitedFunction(s ...PeriodicLimitedFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []Function
func PCMFunctionsToSliceFunction(s ...PCMFunction) []Function {
	out := make([]Function, len(s))
	for i := range out {
		out[i] = s[i].(Function)
	}
	return out
}

// converts []PCMFunction to []LimitedFunction
func PCMFunctionsToSliceLimitedFunction(s ...PCMFunction) []LimitedFunction {
	out := make([]LimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(LimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicLimitedFunction
func PCMFunctionsToSlicePeriodicLimitedFunction(s ...PCMFunction) []PeriodicLimitedFunction {
	out := make([]PeriodicLimitedFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicLimitedFunction)
	}
	return out
}

// converts []PCMFunction to []PeriodicFunction
func PCMFunctionsToSlicePeriodicFunction(s ...PCMFunction) []PeriodicFunction {
	out := make([]PeriodicFunction, len(s))
	for i := range out {
		out[i] = s[i].(PeriodicFunction)
	}
	return out
}/*  Hal3 Fri Mar 4 20:27:57 GMT 2016 go version go1.5.1 linux/amd64
=== RUN   TestNoiseSave
--- PASS: TestNoiseSave (0.90s)
=== RUN   TestSaveLoad
--- PASS: TestSaveLoad (0.01s)
=== RUN   TestSaveWav
--- PASS: TestSaveWav (0.00s)
=== RUN   TestLoad
--- PASS: TestLoad (0.02s)
=== RUN   TestLoadChannels
--- PASS: TestLoadChannels (0.07s)
=== RUN   TestStackPCMs
--- PASS: TestStackPCMs (0.08s)
=== RUN   TestMultiplexTones
--- PASS: TestMultiplexTones (0.04s)
=== RUN   TestSaveLoadSave
--- PASS: TestSaveLoadSave (0.06s)
=== RUN   TestPiping
--- PASS: TestPiping (0.01s)
=== RUN   TestImagingSine
--- PASS: TestImagingSine (0.29s)
=== RUN   TestImaging
--- PASS: TestImaging (0.33s)
=== RUN   TestComposable
--- PASS: TestComposable (1.62s)
=== RUN   TestStackimage
--- PASS: TestStackimage (1.08s)
=== RUN   TestMultiplexImage
--- PASS: TestMultiplexImage (1.02s)
=== RUN   ExampleSquare
--- PASS: ExampleSquare (0.00s)
=== RUN   ExamplePulse
--- PASS: ExamplePulse (0.00s)
=== RUN   ExampleRamp
--- PASS: ExampleRamp (0.00s)
=== RUN   ExampleSine
--- PASS: ExampleSine (0.00s)
=== RUN   ExampleSigmoid
--- PASS: ExampleSigmoid (0.00s)
=== RUN   ExampleADSREnvelope
--- PASS: ExampleADSREnvelope (0.00s)
=== RUN   ExampleReflect
--- PASS: ExampleReflect (0.00s)
=== RUN   ExampleMultiplex
--- PASS: ExampleMultiplex (0.00s)
=== RUN   ExampleStack
--- PASS: ExampleStack (0.00s)
=== RUN   ExampleTrigger
--- PASS: ExampleTrigger (0.00s)
=== RUN   ExampleNoise
--- PASS: ExampleNoise (0.00s)
=== RUN   ExampleBitPulses
--- PASS: ExampleBitPulses (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	5.587s
Fri Mar 4 20:28:04 GMT 2016 */

