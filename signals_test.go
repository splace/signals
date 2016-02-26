package signals

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	s := Square{UnitX}
	for t := x(0); t < 2*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestPulse(t *testing.T) {
	s := Pulse{UnitX}
	for t := x(0); t < 2*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}
func TestRamp(t *testing.T) {
	s := RampUp{UnitX}
	for t := x(0); t < 2*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
	s2 := RampDown{UnitX}
	for t := x(0); t < 2*UnitX; t += UnitX / 10 {
		fmt.Print(s2.Call(t))
	}
	fmt.Println()
}
func TestSine(t *testing.T) {
	s := Sine{UnitX}
	for t := x(0); t < 2*UnitX; t += UnitX / 16 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}
func TestSigmoid(t *testing.T) {
	s := Sigmoid{UnitX}
	for t := x(-2 * UnitX); t < 2*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestADSREnvelope(t *testing.T) {
	s := NewADSREnvelope(UnitX, UnitX, UnitX, Maxy/2, UnitX)
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestReflect(t *testing.T) {
	s := Reflected{NewADSREnvelope(UnitX, UnitX, UnitX, Maxy/2, UnitX)}
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestMulti(t *testing.T) {
	s := Multiplex{Sine{UnitX * 5}, Sine{UnitX * 10}}
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestStack(t *testing.T) {
	s := Stack{Sine{UnitX * 5}, Sine{UnitX * 10}}
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestTrigger(t *testing.T) {
	s := Triggered{NewADSREnvelope(UnitX, UnitX, UnitX, Maxy/2, UnitX), Maxy / 3 * 2, true, UnitX / 100, UnitX * 10, 0, nil, 0, false}
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
	fmt.Println(s.Shift)
	//s.Trigger = Maxy / 3
	s.Rising = false
	for t := x(0); t < 5*UnitX; t += UnitX / 10 {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
	fmt.Println(s.Shift)
	fmt.Println()
}

