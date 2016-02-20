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
	s := Multiplex{Sine{UnitTime * 5}, Sine{UnitTime * 10}}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestStack(t *testing.T) {
	s := Stack{Sine{UnitTime * 5},Sine{UnitTime * 10}}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
}

func TestTrigger(t *testing.T) {
	s := Triggered{NewADSREnvelope(UnitTime, UnitTime, UnitTime, MaxLevel/2, UnitTime), MaxLevel / 3 * 2, true, UnitTime / 100, UnitTime * 10, 0, nil, 0, false}
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	fmt.Println(s.Delay)
	//s.Trigger = MaxLevel / 3
	s.Rising = false
	for t := interval(0); t < 5*UnitTime; t += UnitTime / 10 {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	fmt.Println(s.Delay)
	fmt.Println()
}

func TestNoise(t *testing.T) {
	s := NewNoise()
	for t := interval(0); t < 40*UnitTime; t += UnitTime {
		fmt.Print(s.Level(t))
	}
	fmt.Println()
	var file *os.File
	var err error
	if file, err = os.Create("Noise.wav"); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, s, UnitTime, 8000, 1)
	var file2 *os.File
	if file2, err = os.Create("Noise2.wav"); err != nil {
		panic(err)
	}
	defer file2.Close()
	Encode(file2, s, UnitTime, 16000, 1)
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
	m := NewTone(UnitTime/1000, -6)
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
	m := NewTone(UnitTime/100, -6)

	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, UnitTime, 8000, 1)
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
func TestCombineSounds(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)

	defer stream.Close()
	wavFile, err := os.Create("M1F1-uint8-AFsp-combined.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile,Stack{noise[0], noise[1]}, noise[0].Duration(), 44100, 1)
}
func TestSaveLoadSave(t *testing.T) {
	m := NewTone(UnitTime/1000, -6)
	wavFile, err := os.Create("TestSaveLoad.wav")
	if err != nil {
		panic(err)
	}
	Encode(wavFile,m, UnitTime, 44100, 2)
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
	wavFile, err= os.Create("TestSaveLoadSave.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	noise[0].Encode(wavFile)
}

func TestPiping(t *testing.T) {
	wavFile, err:= os.Create("TestPiping.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	NewPCM(NewTone(UnitTime/200, -6), UnitTime, 8000, 1).Encode(wavFile)
}


