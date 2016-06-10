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
	if file, err = os.Create("./test output/Noise.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, s, unitX, 8000, 1)
	var file2 *os.File
	if file2, err = os.Create("./test output/Noise2.wav"); err != nil {	panic(err)}else{defer file2.Close()}
	Encode(file2, s, unitX, 16000, 1)
}

//func TestSaveLoad(t *testing.T) {
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

func TestSaveWav(t *testing.T) {
	m := Modulated{Sine{unitX/100}, NewConstant(-6)}

	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("./test output/Sine%+v.wav", m)); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, m, unitX, 8000, 3)
}

func TestLoad(t *testing.T) {
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
	if len(noises) != 2 {
		t.Error("Not Double channel", len(noises), stream.Name())
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
	Encode(wavFile, Stack{noise[0], noise[1]}, noise[0].MaxX(), 44100, 1)
}

func TestMultiplexTones(t *testing.T) {
	m := Modulated{Sine{unitX/1000}, NewConstant(-6)}
	m1 := Modulated{Sine{unitX/100}, NewConstant(-6)}
	var file *os.File
	var err error
	if file, err = os.Create("./test output/MultiplexTones.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, Modulated{m, m1}, 1*unitX, 44100, 1)
}

func TestSaveLoadSave(t *testing.T) {
	m := Modulated{Sine{unitX/1000}, NewConstant(-6)}
	wavFile, err := os.Create("./test output/TestSaveLoad.wav")
	if err != nil {
		panic(err)
	}
	Encode(wavFile, m, unitX, 44100, 2)
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

func TestPiping(t *testing.T) {
	wavFile, err := os.Create("./test output/TestPiping.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, Modulated{Sine{unitX/200}, NewConstant(-6)}, unitX, 8000, 1)
}

func TestRawPCM(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/TestRaw.wav"); err != nil {panic(err)}else{defer file.Close()}
	PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})}.Encode(file)
}

func TestSplitPCM(t *testing.T) {
	var wavFileHead *os.File
	var wavFileTail *os.File
	var err error
	if wavFileHead, err = os.Create("./test output/TestSplitHead.wav"); err != nil {panic(err)}else{defer wavFileHead.Close()}
	if wavFileTail, err = os.Create("./test output/TestSplitTail.wav"); err != nil {panic(err)}else{defer wavFileTail.Close()}
	sh, st := PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})}.Split(X(1.01))
	sh.Encode(wavFileHead)
	st.Encode(wavFileTail)
}


func TestEnocdePCMToShortLength(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/EnocdePCMToShortLength.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file,PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})}, unitX, 5, 2)
}


func TestEnocdeShiftedPCM(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/EnocdeShiftedPCM.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file,Shifted{PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})},unitX}, unitX/2, 5, 2)
}



