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


