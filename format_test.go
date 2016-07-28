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


func TestFormatSaveWav(t *testing.T) {
	m := Modulated{Sine{unitX/100}, NewConstant(-6)}

	var file *os.File
	var err error
	for _,bytes:=range([]uint8{1,2,3,4,6}){
		if file, err = os.Create(fmt.Sprintf("./test output/%vbit-Sine%+v.wav",bytes*8, m)); err != nil {panic(err)}else{defer file.Close()}
		Encode(file, bytes, 8000, unitX, m)
	}
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

func TestFormatPCMMultiChannelSave(t *testing.T) {
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

func TestFormatProceduralMultiChannelSave(t *testing.T) {
	wavFile, err := os.Create("./test output/TonesToChannels.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	ss:=[]LimitedSignal{Modulated{Sine{unitX/200}, Pulse{X(1)}},Modulated{Sine{unitX/250}, Pulse{X(1)}}}
	Encode(wavFile, 1, 8000, ss[0].MaxX(), PromoteToSignals(ss)...)
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
	var file *os.File
	var err error
	if file, err = os.Create("./test output/MultiplexTones.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 1, 44100, 1*unitX, Modulated{Sine{unitX/1000}, Sine{unitX/100}})
}

func TestFormatSaveLoadSave(t *testing.T) {
	m := Modulated{Sine{unitX/1000}, NewConstant(-6)}
	wavFile, err := os.Create("./test output/SaveLoad.wav")
	if err != nil {
		panic(err)
	}
	Encode(wavFile, 2, 44100, unitX, m)
	wavFile.Close()
	stream, err := os.Open("./test output/SaveLoad.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	if err != nil {
		panic(err)
	}

	stream.Close()
	wavFile, err = os.Create("./test output/SaveLoadSave.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	EncodeLike(wavFile, noise[0], noise[0])
}

func TestFormatPiping(t *testing.T) {
	wavFile, err := os.Create("./test output/Piping.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, 1, 8000, unitX, Modulated{Sine{unitX/200}, NewConstant(-6)})
}

func TestFormatShortcutEncoding(t *testing.T) {
	file, err := os.Open("middlec.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	noises, err := Decode(file)
	if err != nil {
		panic(err)
	}
	wavFile, err := os.Create("./test output/OffsetPCMShortcut.wav")
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile, 1, 8000, unitX*2, Offset{noises[0], X(1)})
}

