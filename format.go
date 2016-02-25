package signals

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Encode a function as PCM data, one channel, in a Riff wave container.
func Encode(w io.Writer, s Function, length x, sampleRate uint32, sampleBytes uint8) {
	binaryWrite := func(w io.Writer, d interface{}) {
		if err := binary.Write(w, binary.LittleEndian, d); err != nil {
			log.Println("Encode failure:" + err.Error() + fmt.Sprint(w, d))
		}
	}
	samplePeriod := MultiplyX(1/float32(sampleRate), UnitX)
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
				binaryWrite(w, uint8(s.Call(x(i)*samplePeriod)>>(yBits-8)+128))
			}
		}
	case 2:
		if pcm, ok := s.(PCM16bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				binaryWrite(w, int16(s.Call(x(i)*samplePeriod)>>(yBits-16)))
			}
		}
	case 3:
		if pcm, ok := s.(PCM24bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			buf := bytes.NewBuffer(make([]byte, 4))
			for ; i < samples; i++ {
				binaryWrite(buf, int32(s.Call(x(i)*samplePeriod)>>(yBits-32)))
				w.Write(buf.Bytes()[1:])
			}
		}

	case 4:
		if pcm, ok := s.(PCM32bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			w.Write(pcm.data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				binaryWrite(w, int32(s.Call(x(i)*samplePeriod)>>(yBits-32)))
			}
		}
	}
}

// PCMFunction is a Pulse-code modulated Function's behaviour
type PCMFunction interface {
	LimitedFunction
	Periodic
	peaker
	Encode(w io.Writer)
}

// PCM is the state and behaviour common to all PCM.
type PCM struct {
	samplePeriod x
	length       x
	Peak         y
	data         []uint8
}

func (p PCM) Period() x {
	return p.samplePeriod
}

func (p PCM) MaxX() x {
	return p.length
}

func (p PCM) PeakY() y {
	return p.Peak
}

// make a PCMFunction type, from a Function, using particular parameters,
func NewPCM(s Function, length x, sampleRate uint32, sampleBytes uint8) PCMFunction {
	out, in := io.Pipe()
	go func() {
		Encode(in, s, length, sampleRate, sampleBytes)
		in.Close()
	}()
	channels, _ := Decode(out)
	out.Close()
	return channels[0].(PCMFunction)
}

// encode a LimitedFunction with the precision and sampleRate of a given PCMFunction
func EncodeLike(w io.Writer, p LimitedFunction, s PCMFunction) {
	switch f := s.(type) {
	case PCM8bit:
		NewPCM(p, p.MaxX(), uint32(UnitX/f.Period()), 1).Encode(w)
	case PCM16bit:
		NewPCM(p, p.MaxX(), uint32(UnitX/f.Period()), 2).Encode(w)
	case PCM24bit:
		NewPCM(p, p.MaxX(), uint32(UnitX/f.Period()), 3).Encode(w)
	case PCM32bit:
		NewPCM(p, p.MaxX(), uint32(UnitX/f.Period()), 4).Encode(w)
	}
	return
}

// 8 bit PCMFunction
// unlike the other precisions of PCM, that use signed data, 8bit uses un-signed, the default OpenAL and wave file representation.
type PCM8bit struct {
	PCM
}

func (s PCM8bit) Call(offset x) y {
	index := int(offset / s.samplePeriod)
	if index < 0 || index >= len(s.data)-1 {
		return 0
	}
	return s.call(s.data[index])
}

func (s PCM8bit) call(b byte) y {
	return y(b-128) * (Maxy >> 7)
}
func (s PCM8bit) encode(y y) byte {
	return byte(y>>(yBits-8) + 128)
}

func (s PCM8bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(UnitX/s.Period()), 1)
}

// 16 bit PCM Function
type PCM16bit struct {
	PCM
}

func (s PCM16bit) Call(offset x) y {
	index := int(offset/s.samplePeriod) * 2
	if index < 0 || index >= len(s.data)-3 {
		return 0
	}
	return s.call(s.data[index], s.data[index+1])
}

func (s PCM16bit) call(b1, b2 byte) y {
	return y(int16(b1)|int16(b2)<<8) * (Maxy >> 15)
}
func (s PCM16bit) encode(y y) (byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF)
}

func (s PCM16bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(UnitX/s.Period()), 2)
}

// 24 bit PCM Function
type PCM24bit struct {
	PCM
}

func (s PCM24bit) Call(offset x) y {
	index := int(offset/s.samplePeriod) * 3
	if index < 0 || index >= len(s.data)-4 {
		return 0
	}
	return s.call(s.data[index], s.data[index+1], s.data[index+2])
}
func (s PCM24bit) call(b1, b2, b3 byte) y {
	return y(int32(b1)|int32(b2)<<8|int32(b3)<<16) * (Maxy >> 23)
}
func (s PCM24bit) encode(y y) (byte, byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF), byte(y >> (yBits - 24) & 0xFF)
}

func (s PCM24bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(UnitX/s.Period()), 3)
}

// 32 bit PCM Function
type PCM32bit struct {
	PCM
}

func (s PCM32bit) Call(offset x) y {
	index := int(offset/s.samplePeriod) * 4
	if index < 0 || index >= len(s.data)-5 {
		return 0
	}
	return s.call(s.data[index], s.data[index+1], s.data[index+2], s.data[index+3])
}
func (s PCM32bit) call(b1, b2, b3, b4 byte) y {
	return y(int32(b1)|int32(b2)<<8|int32(b3)<<16|int32(b4)<<24) * (Maxy >> 31)
}
func (s PCM32bit) encode(y y) (byte, byte, byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF), byte(y >> (yBits - 24) & 0xFF), byte(y >> (yBits - 32) & 0xFF)
}

func (s PCM32bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(UnitX/s.Period()), 4)
}

type ErrWavParse struct {
	description string
}

func (e ErrWavParse) Error() string {
	return fmt.Sprintf("WAVE Parse,%s", e.description)
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

// Decode a stream into an array of LimitedFunctions.
// one LimitedFunction for each channel in the encoding.
func Decode(wav io.Reader) ([]LimitedFunction, error) {
	var header riffHeader
	var formatHeader chunkHeader
	var format PCMformat
	var dataHeader chunkHeader
	if err := binary.Read(wav, binary.LittleEndian, &header); err != nil {
		return nil, ErrWavParse{"Header not complete."}
	}
	if header.C1 != 'R' || header.C2 != 'I' || header.C3 != 'F' || header.C4 != 'F' || header.C5 != 'W' || header.C6 != 'A' || header.C7 != 'V' || header.C8 != 'E' {
		return nil, ErrWavParse{"Not RIFF/WAVE format."}
	}
	//var runningBytes int =16
	if err := binary.Read(wav, binary.LittleEndian, &formatHeader); err != nil {
		return nil, ErrWavParse{"Chunk incomplete."}
	}
	// TODO skip other chunks
	if formatHeader.C1 != 'f' || formatHeader.C2 != 'm' || formatHeader.C3 != 't' || formatHeader.C4 != ' ' || formatHeader.DataLen != 16 {
		return nil, ErrWavParse{"No format chunk."}
	}

	if err := binary.Read(wav, binary.LittleEndian, &format); err != nil {
		return nil, ErrWavParse{"Format chunk incomplete."}
	}
	if format.Code != 1 {
		return nil, errors.New("only PCM supported.")
	}
	if format.Channels == 0 || format.Channels > 2 {
		return nil, errors.New("only mono or stereo PCM supported.")
	}
	if format.Bits%8 != 0 {
		return nil, ErrWavParse{"not whole byte samples size!"}
	}

	//nice TODO a "LIST" chunk with, 3 fields third being "INFO", can contain "ICOP" and "ICRD" chunks providing copyright and creation date information.

	//	ByteRate    uint32
	//	SampleBytes uint16

	// need to skip any non-"data" chucks
	if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
		return nil, ErrWavParse{"Chunk header incomplete."}
	}
	for dataHeader.C1 != 'd' || dataHeader.C2 != 'a' || dataHeader.C3 != 't' || dataHeader.C4 != 'a' {
		var err error
		if s, ok := wav.(io.Seeker); ok {
			_, err = s.Seek(int64(dataHeader.DataLen), os.SEEK_CUR) // seek relative to current file pointer
		} else {
			_, err = io.CopyN(ioutil.Discard, wav, int64(dataHeader.DataLen))
		}
		if err != nil {
			return nil, ErrWavParse{string(dataHeader.C1) + string(dataHeader.C2) + string(dataHeader.C3) + string(dataHeader.C4) + " " + err.Error()}
		}

		if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
			return nil, ErrWavParse{"Chunk header incomplete."}
		}
	}

	//if dataHeader.DataLen!=header.DataLen-36 {return nil, ErrWavParse{fmt.Sprintf("data chunk size mismatch. %v+36!=%v",dataHeader.DataLen,header.DataLen), []byte(fmt.Sprintf("%#v",dataHeader))}}	//  this is only true for non-extensible wav, ie non-microsoft
	if dataHeader.DataLen%uint32(format.Channels) != 0 {
		return nil, ErrWavParse{"sound sample data length not divisable by channel count"}
	}

	sampleData := make([]byte, dataHeader.DataLen)

	samples := dataHeader.DataLen / uint32(format.Channels) / uint32(format.Bits/8)
	var s uint32
	for ; s < samples; s++ {
		// deinterlace channels by reading directly into separate consecutive blocks
		var c uint32
		for ; c < uint32(format.Channels); c++ {
			if n, err := wav.Read(sampleData[(c*samples+s)*uint32(format.Bits/8) : (c*samples+s+1)*uint32(format.Bits/8)]); err != nil || n != int(format.Bits/8) {
				return nil, ErrWavParse{"data incomplete"}
			}

		}
	}
	functions := make([]LimitedFunction, format.Channels)

	var c uint32
	if format.Bits == 8 {
		for ; c < uint32(format.Channels); c++ {
			functions[c] = PCM8bit{PCM{UnitX / x(format.SampleRate), UnitX / x(format.SampleRate) * x(samples), 0, sampleData[c*samples : (c+1)*samples]}}
		}
	} else if format.Bits == 16 {
		for ; c < uint32(format.Channels); c++ {
			functions[c] = PCM16bit{PCM{UnitX / x(format.SampleRate), UnitX / x(format.SampleRate) * x(samples), 0, sampleData[c*samples*2 : (c+1)*samples*2]}}
		}

	} else if format.Bits == 24 {
		for ; c < uint32(format.Channels); c++ {
			functions[c] = PCM24bit{PCM{UnitX / x(format.SampleRate), UnitX / x(format.SampleRate) * x(samples), 0, sampleData[c*samples*3 : (c+1)*samples*3]}}
		}

	} else if format.Bits == 32 {
		for ; c < uint32(format.Channels); c++ {
			functions[c] = PCM32bit{PCM{UnitX / x(format.SampleRate), UnitX / x(format.SampleRate) * x(samples), 0, sampleData[c*samples*4 : (c+1)*samples*4]}}
		}

	}
	return functions, nil
}
