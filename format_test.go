package signals

import (
	"fmt"
	"os"
	"testing"
)

func TestNoiseSave(t *testing.T) {
	s := NewNoise()
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


func TestSaveLoad(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/multi.gob"); err != nil {
		panic(err)
	}
	m := NewTone(UnitX/1000, -6)
	if err := m.Save(file); err != nil {
		panic("unable to save")
	}
	file.Close()

	if file, err = os.Open("./test output/multi.gob"); err != nil {
		panic(err)
	}
	defer file.Close()

	m1 := Multiplex{}
	if err := m1.Load(file); err != nil {
		panic("unable to load")
	}
	if fmt.Sprintf("%#v", m1)!=fmt.Sprintf("%#v", m) {
		t.Errorf("%#v != %#v", m, m1)
	}

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
	if len(noises)!=1 {
		t.Error("Not Single channel", len(noises),stream.Name())
	}
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
	if len(noises)!=2 {
		t.Error("Not Double channel", len(noises),stream.Name())
	}
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
	stream, err := os.Open("./test output/TestSaveLoad.wav")
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


