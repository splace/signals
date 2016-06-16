package signals

import (
	"fmt"
	"os"
	"testing"
)

func TestFormatNoiseSave(t *testing.T) {
	s := NewNoise()
	var file *os.File
	var err error
	if file, err = os.Create("./test output/Noise.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 1, 8000, unitX, s)
	var file2 *os.File
	if file2, err = os.Create("./test output/Noise2.wav"); err != nil {	panic(err)}else{defer file2.Close()}
	Encode(file2, 1, 16000, unitX, s)
}

//func TestFormatSaveLoad(t *testing.T) {
//	var file *os.File
//	var err error
//	if file, err = os.Create("./test output/multi.gob"); err != nil {
//		panic(err)
//	}
//	m := NewModulated(NewTone(unitX/1000, -6))
//	if err := m.Save(file); err != nil {
//		panic("unable to save")
//	}
//	file.Close()
//
//	if file, err = os.Open("./test output/multi.gob"); err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	m1 := Modulated{}
//	if err := m1.Load(file); err != nil {
//		panic("unable to load")
//	}
//	if fmt.Sprintf("%#v", m1) != fmt.Sprintf("%#v", m) {
//		t.Errorf("%#v != %#v", m, m1)
//	}
//
//}

func TestFormatSaveWav(t *testing.T) {
	m := Modulated{Sine{unitX/100}, NewConstant(-6)}

	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("./test output/Sine%+v.wav", m)); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 3, 8000, unitX, m)
}

func TestFormatLoad(t *testing.T) {
	file, err := os.Open("middlec.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	noises, err := Decode(file)
	if err != nil {
		panic(err)
	}
	if len(noises) != 1 {
		t.Error("Not Single channel", len(noises), file.Name())
	}
}

func TestFormatLoadChannels(t *testing.T) {
	stream, err := os.Open("pcm0808s.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := Decode(stream)
	if err != nil {
		panic(err)
	}
	if len(noises) != 2 {
		t.Error("Not Double channel", len(noises), stream.Name())
	}
}

func TestFormatMultiChannelSave(t *testing.T) {
	stream, err := os.Open("pcm0808s.wav")
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	noises, err := Decode(stream)
	if err != nil {
		panic(err)
	}
	wavFile, err := os.Create("./test output/pcm0808s.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile,4, 44100, noises[0].MaxX(), noises[0],noises[1])
}


func TestFormatStackPCMs(t *testing.T) {
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
	Encode(wavFile, 4, 44100, noise[0].MaxX(), Stack{noise[0], noise[1]})
}

func TestFormatMultiplexTones(t *testing.T) {
	m := Modulated{Sine{unitX/1000}, NewConstant(-6)}
	m1 := Modulated{Sine{unitX/100}, NewConstant(-6)}
	var file *os.File
	var err error
	if file, err = os.Create("./test output/MultiplexTones.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 1, 44100, 1*unitX, Modulated{m, m1})
}

func TestFormatSaveLoadSave(t *testing.T) {
	m := Modulated{Sine{unitX/1000}, NewConstant(-6)}
	wavFile, err := os.Create("./test output/TestSaveLoad.wav")
	if err != nil {
		panic(err)
	}
	Encode(wavFile, 2, 44100, unitX, m)
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
	EncodeLike(wavFile, noise[0], noise[0])
}

func TestFormatPiping(t *testing.T) {
	wavFile, err := os.Create("./test output/TestPiping.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, 1, 8000, unitX, Modulated{Sine{unitX/200}, NewConstant(-6)})
}

