package signals

import (
	"encoding/gob"
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

func TestGob(t *testing.T) {

	file, err := os.Create("signal.gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s:= Sine{UnitTime / 1000}
	//var s Multi
	//s = append(s, Sine{UnitTime / 1000}, Constant{MaxLevel / 2})

	/*	s1,err := json.Marshal(s)
			 	if err != nil {
					panic(err)
				}

				fmt.Fprint(file,string(s1))
		fmt.Fprintf(file, "%+v", s)
		fmt.Fprintf(file, "%#v\n", s)
	*/
	fmt.Fprintf(file, "%v\n", s)
	fmt.Fprintf(file, "%s\n", s)
	fmt.Fprintf(file, "%#v\n", s)
	fmt.Fprintf(file, "%+v\n", s)

	interfaceEncode := func(enc *gob.Encoder, s Signal) {
		err := enc.Encode(&s)
		if err != nil {
			t.Fatal("encode error:", err)
		}
	}

	gob.Register(Sine{})
	//gob.Register(Constant{})
	//gob.Register(Multi{})

	enc := gob.NewEncoder(file)
	interfaceEncode(enc, s)

}

func TestUnmarshal(t *testing.T) {
	file, err := os.Open("signal.gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	interfaceDecode := func(dec *gob.Decoder) Signal {
		var s Signal
		err := dec.Decode(&s)
		if err != nil {
			t.Fatal("decode error:", err)
		}
		return s
	}

	dec := gob.NewDecoder(file)
	s1 := interfaceDecode(dec)

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

	fmt.Fscanf(file, "%#v", &s1)
	fmt.Printf("%#v\n", s1)

}


