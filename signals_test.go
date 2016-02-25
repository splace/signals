package signals

import (
	"fmt"
	"math/big"
	"os"
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

func TestNoise(t *testing.T) {
	s := NewNoise()
	for t := x(0); t < 40*UnitX; t += UnitX {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
	var file *os.File
	var err error
	if file, err = os.Create("./test output/Noise.wav"); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, s, UnitX, 8000, 1)
	var file2 *os.File
	if file2, err = os.Create("./test output/Noise2.wav"); err != nil {
		panic(err)
	}
	defer file2.Close()
	Encode(file2, s, UnitX, 16000, 1)
}

func TestBitPulses(t *testing.T) {
	i := new(big.Int)
	_, err := fmt.Sscanf("01110111011101110111011101110111", "%b", i)
	if err != nil {
		panic(i)
	}
	s := PulsePattern{*i, UnitX}
	for t := x(0); t < 40*UnitX; t += UnitX {
		fmt.Print(s.Call(t))
	}
	fmt.Println()
}

func TestSaveLoad(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/multi.gob"); err != nil {
		panic(err)
	}
	//m:=Multi{Sine{UnitTime / 1000},Constant{50}}
	m := NewTone(UnitX/1000, -6)
	if err := m.Save(file); err != nil {
		panic("unable to save")
	}
	file.Close()

	if file, err = os.Open("multi.gob"); err != nil {
		panic(err)
	}
	defer file.Close()

	m1 := Multiplex{}
	if err := m1.Load(file); err != nil {
		panic("unable to load")
	}
	fmt.Printf("%#v\n", m1)

}

func TestSaveWav(t *testing.T) {
	m := NewTone(UnitX/100, -6)

	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, UnitX, 8000, 1)
}
func TestLoad(t *testing.T) {
	stream, err := os.Open("middlec.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := Decode(stream)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(noises))
}

func TestLoadChannels(t *testing.T) {
	stream, err := os.Open("pcm0808s.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := Decode(stream)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(noises))
}
func TestStackPCMs(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)

	defer stream.Close()
	wavFile, err := os.Create("./test output/StackPCMs.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, Stack{noise[0], noise[1]}, noise[0].(PCMFunction).MaxX(), 44100, 1)
}
func TestMultiplexTones(t *testing.T) {
	m := NewTone(UnitX/1000, -6)
	m1 := NewTone(UnitX/100, -6)
	wavFile, err := os.Create("./test output/MultiplexTones.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, Multiplex{m, m1}, 1*UnitX, 44100, 1)
}
func TestSaveLoadSave(t *testing.T) {
	m := NewTone(UnitX/1000, -6)
	wavFile, err := os.Create("./test output/TestSaveLoad.wav")
	if err != nil {
		panic(err)
	}
	Encode(wavFile, m, UnitX, 44100, 2)
	wavFile.Close()
	stream, err := os.Open("TestSaveLoad.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	if err != nil {
		panic(err)
	}

	stream.Close()
	wavFile, err = os.Create("./test output/TestSaveLoadSave.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	noise[0].(PCMFunction).Encode(wavFile)
}

func TestPiping(t *testing.T) {
	wavFile, err := os.Create("./test output/TestPiping.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	NewPCM(NewTone(UnitX/200, -6), UnitX, 8000, 1).Encode(wavFile)
}
func TestImagingSine(t *testing.T) {
	Save("./test output/Sine",NewFunctionImage(Multiplex{Sine{UnitX}, Pulse{UnitX}},800,300))
}
func TestImaging(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)

	defer stream.Close()
//	Save("./test output/M1F1-uint8-AFsp.wav",NewFunctionImage(Multiplex{noise[0], Pulse{UnitX*4}},3200,300))   // first second
	Save("./test output/M1F1-uint8-AFsp.wav",NewFunctionImage(noise[0],int(noise[0].MaxX()/UnitX*800),300))    // 800 pixels per second width
}



/*  Hal3 Thu Feb 25 22:25:40 GMT 2016 go version go1.5.1 linux/amd64
FAIL	_/home/simon/Dropbox/github/working/signals [build failed]
Thu Feb 25 22:25:40 GMT 2016 */
/*  Hal3 Thu Feb 25 22:26:13 GMT 2016 go version go1.5.1 linux/amd64
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
=== RUN   TestMulti
     0.00%     0.79%     3.12%     6.90%    11.98%    18.16%    25.20%    32.81%    40.68%    48.48%    55.90%    62.61%    68.32%    72.75%    75.69%    76.94%    76.40%    73.99%    69.72%    63.65%    55.90%    46.66%    36.16%    24.67%    12.51%     0.00%   -12.51%   -24.67%   -36.16%   -46.66%   -55.90%   -63.65%   -69.72%   -73.99%   -76.40%   -76.94%   -75.69%   -72.75%   -68.32%   -62.61%   -55.90%   -48.48%   -40.68%   -32.81%   -25.20%   -18.16%   -11.98%    -6.90%    -3.12%    -0.79%
--- PASS: TestMulti (0.00s)
=== RUN   TestStack
     0.00%     9.41%    18.70%    27.78%    36.52%    44.84%    52.63%    59.81%    66.30%    72.03%    76.94%    80.99%    84.13%    86.35%    87.64%    88.00%    87.46%    86.03%    83.77%    80.72%    76.94%    72.52%    67.52%    62.04%    56.17%    50.00%    43.63%    37.17%    30.71%    24.34%    18.16%    12.26%     6.72%     1.60%    -3.02%    -7.10%   -10.59%   -13.45%   -15.67%   -17.24%   -18.16%   -18.45%   -18.13%   -17.24%   -15.82%   -13.94%   -11.65%    -9.04%    -6.17%    -3.13%
--- PASS: TestStack (0.00s)
=== RUN   TestTrigger
    67.00%    77.00%    87.00%    97.00%    96.50%    91.50%    86.50%    81.50%    76.50%    71.50%    66.50%    61.50%    56.50%    51.50%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    46.50%    41.50%    36.50%    31.50%    26.50%    21.50%    16.50%    11.50%     6.50%     1.50%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
     0.67s
    66.50%    61.50%    56.50%    51.50%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    50.00%    46.50%    41.50%    36.50%    31.50%    26.50%    21.50%    16.50%    11.50%     6.50%     1.50%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
     1.67s

--- PASS: TestTrigger (0.00s)
=== RUN   TestNoise
    23.94%   -52.49%     8.21%    -9.87%   -74.46%   -68.54%   -31.13%   -28.89%    11.03%    43.01%   -71.97%   -35.88%   -58.86%    47.80%    21.68%   -34.58%   -66.41%    10.38%     4.28%   -14.14%   -17.82%   -31.24%    22.84%   -21.90%    17.72%    23.27%    38.15%    65.67%   -72.58%   -66.54%   -33.93%     4.60%   -42.08%   -36.43%   -48.60%   -10.65%   -17.75%    25.50%    23.76%   -87.69%
--- PASS: TestNoise (0.97s)
=== RUN   TestBitPulses
   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%   100.00%   100.00%   100.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%     0.00%
--- PASS: TestBitPulses (0.00s)
=== RUN   TestSaveLoad
signals.Multiplex{signals.Sine{Cycle:1000000}, signals.Constant{Constant:4611686018427387392}}
--- PASS: TestSaveLoad (0.01s)
=== RUN   TestSaveWav
--- PASS: TestSaveWav (0.05s)
=== RUN   TestLoad
1
--- PASS: TestLoad (0.02s)
=== RUN   TestLoadChannels
2
--- PASS: TestLoadChannels (0.07s)
=== RUN   TestStackPCMs
--- PASS: TestStackPCMs (0.50s)
=== RUN   TestMultiplexTones
--- PASS: TestMultiplexTones (0.22s)
=== RUN   TestSaveLoadSave
--- PASS: TestSaveLoadSave (0.21s)
=== RUN   TestPiping
--- PASS: TestPiping (0.03s)
=== RUN   TestImagingSine
--- PASS: TestImagingSine (0.34s)
=== RUN   TestImaging
--- PASS: TestImaging (0.52s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	2.985s
Thu Feb 25 22:26:17 GMT 2016 */

