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

/*  Hal3 Wed Feb 10 21:17:32 GMT 2016 go version go1.5.1 linux/amd64
=== RUN   TestSquare
   100.00%   100.00%   100.00%   100.00%   100.00%  -100.00%  -100.00%  -100.00%  -100.00%  -100.00%   100.00%   100.00%   100.00%   100.00%   100.00%  -100.00%  -100.00%  -100.00%  -100.00%  -100.00%
--- PASS: TestSquare (0.00s)
=== RUN   TestPulse
   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
--- PASS: TestPulse (0.00s)
=== RUN   TestRamp
     0.00%    10.00%    20.00%    30.00%    40.00%    50.00%    60.00%    70.00%    80.00%    90.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%
   100.00%    90.00%    80.00%    70.00%    60.00%    50.00%    40.00%    30.00%    20.00%    10.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
--- PASS: TestRamp (0.00s)
=== RUN   TestSine
     0.00%    38.27%    70.71%    92.39%   100.00%    92.39%    70.71%    38.27%     0.00%   -38.27%   -70.71%   -92.39%  -100.00%   -92.39%   -70.71%   -38.27%    -0.00%    38.27%    70.71%    92.39%   100.00%    92.39%    70.71%    38.27%     0.00%   -38.27%   -70.71%   -92.39%  -100.00%   -92.39%   -70.71%   -38.27%
--- PASS: TestSine (0.00s)
=== RUN   TestSigmoid
    11.92%    13.01%    14.19%    15.45%    16.80%    18.24%    19.78%    21.42%    23.15%    24.97%    26.89%    28.91%    31.00%    33.18%    35.43%    37.75%    40.13%    42.56%    45.02%    47.50%    50.00%    52.50%    54.98%    57.44%    59.87%    62.25%    64.57%    66.82%    69.00%    71.09%    73.11%    75.03%    76.85%    78.58%    80.22%    81.76%    83.20%    84.55%    85.81%    86.99%
--- PASS: TestSigmoid (0.00s)
=== RUN   TestADSREnvelope
     0.00%    10.00%    20.00%    30.00%    40.00%    50.00%    60.00%    70.00%    80.00%    90.00%   100.00%    95.00%    90.00%    85.00%    80.00%    75.00%    70.00%    65.00%    60.00%    55.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    45.00%    40.00%    35.00%    30.00%    25.00%    20.00%    15.00%    10.00%     5.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
--- PASS: TestADSREnvelope (0.00s)
=== RUN   TestReflect
   100.00%    90.00%    80.00%    70.00%    60.00%    50.00%    40.00%    30.00%    20.00%    10.00%     0.00%     5.00%    10.00%    15.00%    20.00%    25.00%    30.00%    35.00%    40.00%    45.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    55.00%    60.00%    65.00%    70.00%    75.00%    80.00%    85.00%    90.00%    95.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%   100.00%
--- PASS: TestReflect (0.00s)
=== RUN   TestBitPulses
   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
--- PASS: TestBitPulses (0.00s)
=== RUN   TestMarshal
--- PASS: TestMarshal (0.02s)
=== RUN   TestUnmarshal
signals.Product(nil)
--- PASS: TestUnmarshal (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	0.032s
Wed Feb 10 21:17:34 GMT 2016 */

