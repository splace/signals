package signals

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)

func TestSquare(t *testing.T) {
	s := Square{UnitTime}
	for t := Interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestPulse(t *testing.T) {
	s := Pulse{UnitTime}
	for t := Interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}
func TestRamp(t *testing.T) {
	s := RampUp{UnitTime}
	for t := Interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	s2 := RampDown{UnitTime}
	for t := Interval(0); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s2.Level(t))
	}
	fmt.Println()
}
func TestSine(t *testing.T) {
	s := Sine{UnitTime}
	for t := Interval(0); t < 2*UnitTime; t += UnitTime / 16 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}
func TestSigmoid(t *testing.T) {
	s := Sigmoid{UnitTime}
	for t := Interval(-2 * UnitTime); t < 2*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestADSREnvelope(t *testing.T) {
	s := NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime)
	for t := Interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestReflect(t *testing.T) {
	s := Reflected{NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime)}
	for t := Interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestBitPulses(t *testing.T) {
	i := new(big.Int)
	_, err := fmt.Sscanf("01110111011101110111011101110111", "%b", i)
	if err != nil {
		panic(i)
	}
	s := PulsePattern{*i, UnitTime}
	for t := Interval(0); t < 40*UnitTime; t += UnitTime {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestMarshal(t *testing.T) {

	file, err := os.Create("main.js")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var s Product
	s = append(s, Sine{UnitTime / 1000}, Constant{MaxLevel / 2})
	/*	s1,err := json.Marshal(s)
		 	if err != nil {
				panic(err)
			}

			fmt.Fprint(file,string(s1))
	*/
	fmt.Fprintf(file, "%+v", s)
}

func TestUnmarshal(t *testing.T) {
	file, err := os.Open("main.js")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	/*
		s,err:=ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
	*/
	//s:=[]byte("{\"Item\":\"Constant\"}")
	//var s1 struct{Item string}

	/*
				var s1 Constant
			err = json.Unmarshal(s,&s1)
		 	if err != nil {
				panic(err)
			}
	*/
	var s1 Product

	fmt.Fscanf(file, "%#v", &s1)
	fmt.Printf("%#v\n", s1)
}


