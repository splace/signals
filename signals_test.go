package signals

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)

func TestSquare(t *testing.T) {
	s := Square{UnitTime}
	for t := interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestPulse(t *testing.T) {
	s := Pulse{UnitTime}
	for t := interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}
func TestRamp(t *testing.T) {
	s := RampUp{UnitTime}
	for t := interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	s2 := RampDown{UnitTime}
	for t := interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s2.Level(t))
	}
	fmt.Println()
}
func TestSine(t *testing.T) {
	s := Sine{UnitTime}
	for t := interval(0); t < 2*UnitTime; t += UnitTime / 16 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}
func TestSigmoid(t *testing.T) {
	s := Sigmoid{UnitTime}
	for t := interval(-2 * UnitTime); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestADSREnvelope(t *testing.T) {
	s := NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime)
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestReflect(t *testing.T) {
	s := Reflected{NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime)}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestMulti(t *testing.T) {
	s := Multi{Sine{UnitTime * 5}, Sine{UnitTime * 10}}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestSum(t *testing.T) {
	s := Sum{Multi{Sine{UnitTime * 5}, NewConstant(50)}, Multi{Sine{UnitTime * 10}, NewConstant(50)}}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestTrigger(t *testing.T) {
	s := Triggered{NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime), MaxLevel / 3 * 2, true,UnitTime / 100, UnitTime * 10, 0, nil, 0}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	fmt.Println(s.Delay)
	s.Trigger = MaxLevel / 3
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	fmt.Println(s.Delay)
	fmt.Println()
}

func TestBitPulses(t *testing.T) {
	i := new(big.Int)
	_, err := fmt.Sscanf("01110111011101110111011101110111", "%b", i)
	if err != nil {
		panic(i)
	}
	s := PulsePattern{*i, UnitTime}
	for t := interval(0); t < 40*UnitTime; t += UnitTime {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestSaveLoad(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("multi.gob"); err != nil {
		panic(err)
	}
	//m:=Multi{Sine{UnitTime / 1000},Constant{50}}
	m := NewTone(UnitTime/1000, 50)
	if err := m.Save(file); err != nil {
		panic("unable to save")
	}
	file.Close()

	if file, err = os.Open("multi.gob"); err != nil {
		panic(err)
	}
	defer file.Close()

	m1 := Multi{}
	if err := m1.Load(file); err != nil {
		panic("unable to load")
	}
	fmt.Printf("%#v\n", m1)

}

func TestSaveWav(t *testing.T) {
	m := NewTone(UnitTime/100, 50)

	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, UnitTime, 8000, 1)
}

