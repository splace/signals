package signals

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Encode a signal as PCM data, one channel, in a Riff wave container.
func Encode(w io.Writer, s Signal, length interval, sampleRate uint32, sampleBytes uint8) {
	binaryWrite := func(w io.Writer, d interface{}) {
		if err := binary.Write(w, binary.LittleEndian, d); err != nil {
			log.Println("Encode failure:" + err.Error() + fmt.Sprint(w, d))
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
		if pcm, ok := s.(PCM8bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				binaryWrite(w, uint8(s.Level(interval(i)*samplePeriod)>>(LevelBits-8)+128))
			}
		}
	case 2:
		if pcm, ok := s.(PCM16bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				binaryWrite(w, int16(s.Level(interval(i)*samplePeriod)>>(LevelBits-16)))
			}
		}
	case 3:
		if pcm, ok := s.(PCM24bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			buf := bytes.NewBuffer(make([]byte, 4))
			for ; i < samples; i++ {
				binaryWrite(buf, int32(s.Level(interval(i)*samplePeriod)>>(LevelBits-32)))
				w.Write(buf.Bytes()[1:])
			}
		}

	case 4:
		if pcm, ok := s.(PCM32bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				binaryWrite(w, int32(s.Level(interval(i)*samplePeriod)>>(LevelBits-32)))
			}
		}
	}
}

type ErrWavParse struct {
	description string
	source      io.Reader
}

func (e ErrWavParse) Error() string {
	switch st := e.source.(type) {
	case *os.File:
		return fmt.Sprintf("WAVE Parse,%s File:%s", e.description, st.Name())
	}
	return fmt.Sprintf("WAVE Parse,%s:%v", e.description, e.source)
}

// PCMSignal is a Pulse-code modulated Signal's behaviour
type PCMSignal interface {
	LimitedSignal
	ModulatePeriod() interval
	Encode(w io.Writer)
}

// PCM is the state and behaviour common to all PCMSignals.
type PCM struct {
	samplePeriod interval
	length       interval
	data         []uint8
}

func (p PCM) ModulatePeriod() interval {
	return p.samplePeriod
}

func (p PCM) Duration() interval {
	return p.length
}

// make a PCMSignal type, from a Signal, using particular parameters,
func NewPCM(s Signal, length interval, sampleRate uint32, sampleBytes uint8) PCMSignal {
	in, out := io.Pipe()
	go func() {
		Encode(out, s, length, sampleRate, sampleBytes)
		out.Close()
	}()
	noise, _ := Decode(in)
	in.Close()
	return noise[0]
}

// 8 bit PCMSignal
// unlike the other precisions of PCM, that use signed data, 8bit uses un-signed, the default OpenAL and wave file representation.
type PCM8bit struct {
	PCM
}

func (s PCM8bit) Level(offset interval) level {
	index := int(offset / s.samplePeriod)
	if index < 0 || index >= len(s.data)-1 {
		return 0
	}
	return level(s.data[index]-128) * (MaxLevel >> 7)
}
func (s PCM8bit) Encode(w io.Writer) {
	Encode(w, s, s.Duration(), uint32(UnitTime/s.ModulatePeriod()), 1)
}

// 16 bit PCM Signal
type PCM16bit struct {
	PCM
}

func (s PCM16bit) Level(offset interval) level {
	index := int(offset/s.samplePeriod) * 2
	if index < 0 || index >= len(s.data)-3 {
		return 0
	}
	return level(int16(s.data[index])|int16(s.data[index+1])<<8) * (MaxLevel >> 15)
}
func (s PCM16bit) Encode(w io.Writer) {
	Encode(w, s, s.Duration(), uint32(UnitTime/s.ModulatePeriod()), 2)
}

// 24 bit PCM Signal
type PCM24bit struct {
	PCM
}

func (s PCM24bit) Level(offset interval) level {
	index := int(offset/s.samplePeriod) * 3
	if index < 0 || index >= len(s.data)-4 {
		return 0
	}
	return level(int32(s.data[index])|int32(s.data[index+1])<<8|int32(s.data[index+2])<<16) * (MaxLevel >> 23)
}
func (s PCM24bit) Encode(w io.Writer) {
	Encode(w, s, s.Duration(), uint32(UnitTime/s.ModulatePeriod()), 3)
}

// 32 bit PCM Signal
type PCM32bit struct {
	PCM
}

func (s PCM32bit) Level(offset interval) level {
	index := int(offset/s.samplePeriod) * 4
	if index < 0 || index >= len(s.data)-5 {
		return 0
	}
	return level(int32(s.data[index])|int32(s.data[index+1])<<8|int32(s.data[index+2])<<16|int32(s.data[index+3])<<24) * (MaxLevel >> 31)
}
func (s PCM32bit) Encode(w io.Writer) {
	Encode(w, s, s.Duration(), uint32(UnitTime/s.ModulatePeriod()), 4)
}

// RIFF file header holder
type riffHeader struct {
	C1, C2, C3, C4 byte
	DataLen        uint32
	C5, C6, C7, C8 byte
}

// RIFF chunk header holder
type chunkHeader struct {
	C1, C2, C3, C4 byte
	DataLen        uint32
}

// PCM format holder
type PCMformat struct {
	Code        uint16
	Channels    uint16
	SampleRate  uint32
	ByteRate    uint32
	SampleBytes uint16
	Bits        uint16
}

// Decode a stream into an array of PCMSignals.
// one Signal for each channel in the encoding.
func Decode(wav io.Reader) ([]PCMSignal, error) {
	var header riffHeader
	var formatHeader chunkHeader
	var format PCMformat
	var dataHeader chunkHeader
	if err := binary.Read(wav, binary.LittleEndian, &header); err != nil {
		return nil, ErrWavParse{"Header not complete.", wav}
	}
	if header.C1 != 'R' || header.C2 != 'I' || header.C3 != 'F' || header.C4 != 'F' || header.C5 != 'W' || header.C6 != 'A' || header.C7 != 'V' || header.C8 != 'E' {
		return nil, ErrWavParse{"Not RIFF/WAVE format.", wav}
	}
	//var runningBytes int =16
	if err := binary.Read(wav, binary.LittleEndian, &formatHeader); err != nil {
		return nil, ErrWavParse{"Chunk incomplete.", wav}
	}
	// TODO skip other chunks
	if formatHeader.C1 != 'f' || formatHeader.C2 != 'm' || formatHeader.C3 != 't' || formatHeader.C4 != ' ' || formatHeader.DataLen != 16 {
		return nil, ErrWavParse{"No format chunk.", wav}
	}

	if err := binary.Read(wav, binary.LittleEndian, &format); err != nil {
		return nil, ErrWavParse{"Format chunk incomplete.", wav}
	}
	if format.Code != 1 {
		return nil, ErrWavParse{"not PCM format.", wav}
	}
	if format.Channels == 0 || format.Channels > 2 {
		return nil, ErrWavParse{"not mono or stereo.", wav}
	}
	if format.Bits%8 != 0 {
		return nil, ErrWavParse{"not whole byte samples size!", wav}
	}

	//nice TODO a "LIST" chunk with, 3 fields third being "INFO", can contain "ICOP" and "ICRD" chunks providing copyright and creation date information.

	//	ByteRate    uint32
	//	SampleBytes uint16

	// need to skip any non-"data" chucks
	if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
		return nil, ErrWavParse{"Chunk header incomplete.", wav}
	}
	for dataHeader.C1 != 'd' || dataHeader.C2 != 'a' || dataHeader.C3 != 't' || dataHeader.C4 != 'a' {
		var err error
		if s, ok := wav.(io.Seeker); ok {
			_, err = s.Seek(int64(dataHeader.DataLen), os.SEEK_CUR) // seek relative to current file pointer
		} else {
			_, err = io.CopyN(ioutil.Discard, wav, int64(dataHeader.DataLen))
		}
		if err != nil {
			return nil, ErrWavParse{string(dataHeader.C1) + string(dataHeader.C2) + string(dataHeader.C3) + string(dataHeader.C4) + " " + err.Error(), wav}
		}

		if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
			return nil, ErrWavParse{"Chunk header incomplete.", wav}
		}
	}

	//if dataHeader.DataLen!=header.DataLen-36 {return nil, ErrWavParse{fmt.Sprintf("data chunk size mismatch. %v+36!=%v",dataHeader.DataLen,header.DataLen), []byte(fmt.Sprintf("%#v",dataHeader))}}	//  this is only true for non-extensible wav, ie non-microsoft
	if dataHeader.DataLen%uint32(format.Channels) != 0 {
		return nil, ErrWavParse{"sound sample data length not divisable by channel count", wav}
	}

	sampleData := make([]byte, dataHeader.DataLen)

	samples := dataHeader.DataLen / uint32(format.Channels) / uint32(format.Bits/8)
	var s uint32
	for ; s < samples; s++ {
		// deinterlace channels by reading directly into separate blocks
		var c uint32
		for ; c < uint32(format.Channels); c++ {
			if n, err := wav.Read(sampleData[(c*samples+s)*uint32(format.Bits/8) : (c*samples+s+1)*uint32(format.Bits/8)]); err != nil || n != int(format.Bits/8) {
				return nil, ErrWavParse{"data incomplete", wav}
			}

		}
	}
	signals := make([]PCMSignal, format.Channels)

	var c uint32
	if format.Bits == 8 {
		for ; c < uint32(format.Channels); c++ {
			signals[c] = PCM8bit{PCM{UnitTime / interval(format.SampleRate), UnitTime / interval(format.SampleRate) * interval(samples), sampleData[c*samples : (c+1)*samples]}}
		}
	} else if format.Bits == 16 {
		for ; c < uint32(format.Channels); c++ {
			signals[c] = PCM16bit{PCM{UnitTime / interval(format.SampleRate), UnitTime / interval(format.SampleRate) * interval(samples), sampleData[c*samples*2 : (c+1)*samples*2]}}
		}

	} else if format.Bits == 24 {
		for ; c < uint32(format.Channels); c++ {
			signals[c] = PCM24bit{PCM{UnitTime / interval(format.SampleRate), UnitTime / interval(format.SampleRate) * interval(samples), sampleData[c*samples*3 : (c+1)*samples*3]}}
		}

	} else if format.Bits == 32 {
		for ; c < uint32(format.Channels); c++ {
			signals[c] = PCM32bit{PCM{UnitTime / interval(format.SampleRate), UnitTime / interval(format.SampleRate) * interval(samples), sampleData[c*samples*4 : (c+1)*samples*4]}}
		}

	}
	return signals, nil
}
