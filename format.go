package signals

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// encode a signal as PCM data in a Riff wave container (mono wav file format)
func Encode(w io.Writer, s Signal, length interval, sampleRate uint32, sampleBytes uint8) {
	binaryWrite := func(w io.Writer, d interface{}) {
		if err := binary.Write(w, binary.LittleEndian, d); err != nil {
			panic(err)
		}
	}
	samplePeriod := MultiplyInterval(1/float32(sampleRate), UnitTime)
	samples := uint32(length/samplePeriod) + 1
	fmt.Fprint(w, "RIFF")
	binaryWrite(w, samples*uint32(sampleBytes)+36)
	fmt.Fprint(w, "WAVE")
	fmt.Fprint(w, "fmt ")
	binaryWrite(w, uint32(16))
	binaryWrite(w, uint16(1))
	binaryWrite(w, uint16(1))
	binaryWrite(w, sampleRate)
	binaryWrite(w, sampleRate*uint32(sampleBytes))
	binaryWrite(w, uint16(sampleBytes))
	binaryWrite(w, uint16(8*sampleBytes))
	fmt.Fprint(w, "data")
	binaryWrite(w, samples*uint32(sampleBytes))
	var i uint32
	switch sampleBytes {
	case 1:
		for ; i < samples; i++ {
			binaryWrite(w, uint8(s.Level(interval(i)*samplePeriod)>>(LevelBits-8)+128))
		}
	case 2:
		for ; i < samples; i++ {
			binaryWrite(w, int16(s.Level(interval(i)*samplePeriod)>>(LevelBits-16)))
		}
	case 3:
		buf := bytes.NewBuffer(make([]byte, 4))
		for ; i < samples; i++ {
			binaryWrite(buf, int32(s.Level(interval(i)*samplePeriod)>>(LevelBits-32)))
			w.Write(buf.Bytes()[1:])
		}

	case 4:
		for ; i < samples; i++ {
			binaryWrite(w, int32(s.Level(interval(i)*samplePeriod)>>(LevelBits-32)))
		}
	}
}
