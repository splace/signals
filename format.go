package signals

import (
	"bufio"
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
	var err error
	var i uint32
	buf := bufio.NewWriter(w)
	samplePeriod := X(1 / float32(sampleRate))
	samples := uint32(length/samplePeriod) + 1
	writeHeader(buf, sampleRate, samples,  sampleBytes)
	switch sampleBytes {
	case 1:
		if pcm, ok := s.(PCM8bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			buf.Write(pcm.Data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				err = buf.WriteByte(PCM8bitEncode(s.property(x(i) * samplePeriod)))
				if err != nil {break}
			}
		}
	case 2:
		if pcm, ok := s.(PCM16bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			buf.Write(pcm.Data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				b1, b2 := PCM16bitEncode(s.property(x(i) * samplePeriod))
				err = buf.WriteByte(b2)
				err = buf.WriteByte(b1)
				if err != nil {break}
			}
		}
	case 3:
		if pcm, ok := s.(PCM24bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			buf.Write(pcm.Data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				b1, b2, b3 := PCM24bitEncode(s.property(x(i) * samplePeriod))
				err = buf.WriteByte(b3)
				err = buf.WriteByte(b2)
				err = buf.WriteByte(b1)
				if err != nil {break}
			}
		}
	case 4:
		if pcm, ok := s.(PCM32bit); ok && pcm.samplePeriod == samplePeriod && pcm.length == length {
			buf.Write(pcm.Data) // TODO can cope with shorter length
		} else {
			for ; i < samples; i++ {
				b1, b2, b3, b4 := PCM32bitEncode(s.property(x(i) * samplePeriod))
				err = buf.WriteByte(b4)
				err = buf.WriteByte(b3)
				err = buf.WriteByte(b2)
				err = buf.WriteByte(b1)
				if err != nil {break}
			}
		}
	}
	if err != nil {
		log.Println("Encode failure:" + err.Error() + fmt.Sprint(buf))
	}else{
		buf.Flush()
	}
}


func writeHeader(w *bufio.Writer, sampleRate uint32, samples uint32, sampleBytes uint8) {
	binaryWrite := func(w io.Writer, d interface{}) {
		if err := binary.Write(w, binary.LittleEndian, d); err != nil {
			log.Println("Encode failure:" + err.Error() + fmt.Sprint(w, d))
		}
	}
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
	return
}

// PCMFunction is a Pulse-code modulated Function's behaviour
type PCMFunction interface {
	PeriodicLimitedFunction
	Encode(w io.Writer)
}

// make a PCMFunction type, from a Function, using particular parameters,
func NewPCMFunction(s Function, length x, sampleRate uint32, sampleBytes uint8) PCMFunction {
	out, in := io.Pipe()
	go func() {
		Encode(in, s, length, sampleRate, sampleBytes)
		in.Close()
	}()
	channels, _ := Decode(out)
	out.Close()
	return channels[0].(PCMFunction)
}

// PCM is the state and behaviour common to all PCM. Its not a Function, specific PCM<<precison>> types embed this, and then are LimitedPeriodicFunction's.
type PCM struct {
	samplePeriod x
	length       x
	Data         []byte
}

// make a PCM type, from raw bytes.
func NewPCM(sampleRate uint32, sampleBytes uint8, Data []byte) PCM {
	period := X(1 / float32(sampleRate))
	if len(Data)%int(sampleBytes) != 0 {
		log.Println("Byte array not whole number of samples")
	}
	return PCM{period, period * x(len(Data)/int(sampleBytes)), Data}
}

func (p PCM) Period() x {
	return p.samplePeriod
}

func (p PCM) MaxX() x {
	return p.length
}

// encode a LimitedFunction with a sampleRate equal to the Period() of a given PeriodicLimitedFunction, and its precision if its a PCM type, otherwise defaults to 16bit.
func EncodeLike(w io.Writer, p LimitedFunction, s PeriodicLimitedFunction) {
	switch f := s.(type) {
	case PCM8bit:
		NewPCMFunction(p, p.MaxX(), uint32(unitX/f.Period()), 1).Encode(w)
	case PCM16bit:
		NewPCMFunction(p, p.MaxX(), uint32(unitX/f.Period()), 2).Encode(w)
	case PCM24bit:
		NewPCMFunction(p, p.MaxX(), uint32(unitX/f.Period()), 3).Encode(w)
	case PCM32bit:
		NewPCMFunction(p, p.MaxX(), uint32(unitX/f.Period()), 4).Encode(w)
	default:
		NewPCMFunction(p, p.MaxX(), uint32(unitX/f.Period()), 2).Encode(w)
	}
	return
}

// 8 bit PCMFunction.
// (unlike the other precisions of PCM, that use signed data, 8bit uses un-signed, the default OpenAL and wave file representation for 8bit precision.)
type PCM8bit struct {
	PCM
}

func (s PCM8bit) property(offset x) y {
	index := int(offset / s.samplePeriod)
	if index < 0 || index >= len(s.Data)-1 {
		return 0
	}
	return PCM8bitDecode(s.Data[index])
}

func PCM8bitDecode(b byte) y {
	return y(b-128) * (unitY >> 7)
}
func PCM8bitEncode(y y) byte {
	return byte(y>>(yBits-8) + 128)
}

func (s PCM8bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(unitX/s.Period()), 1)
}

// 16 bit PCM Function
type PCM16bit struct {
	PCM
}

func (s PCM16bit) property(offset x) y {
	index := int(offset/s.samplePeriod) * 2
	if index < 0 || index >= len(s.Data)-3 {
		return 0
	}
	return PCM16bitDecode(s.Data[index], s.Data[index+1])
}

func PCM16bitDecode(b1, b2 byte) y {
	return y(int16(b1)|int16(b2)<<8) * (unitY >> 15)
}
func PCM16bitEncode(y y) (byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF)
}

func (s PCM16bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(unitX/s.Period()), 2)
}

// 24 bit PCM Function
type PCM24bit struct {
	PCM
}

func (s PCM24bit) property(offset x) y {
	index := int(offset/s.samplePeriod) * 3
	if index < 0 || index >= len(s.Data)-4 {
		return 0
	}
	return PCM24bitDecode(s.Data[index], s.Data[index+1], s.Data[index+2])
}
func PCM24bitDecode(b1, b2, b3 byte) y {
	return y(int32(b1)|int32(b2)<<8|int32(b3)<<16) * (unitY >> 23)
}
func PCM24bitEncode(y y) (byte, byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF), byte(y >> (yBits - 24) & 0xFF)
}

func (s PCM24bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(unitX/s.Period()), 3)
}

// 32 bit PCM Function
type PCM32bit struct {
	PCM
}

func (s PCM32bit) property(offset x) y {
	index := int(offset/s.samplePeriod) * 4
	if index < 0 || index >= len(s.Data)-5 {
		return 0
	}
	return PCM32bitDecode(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3])
}
func PCM32bitDecode(b1, b2, b3, b4 byte) y {
	return y(int32(b1)|int32(b2)<<8|int32(b3)<<16|int32(b4)<<24) * (unitY >> 31)
}
func PCM32bitEncode(y y) (byte, byte, byte, byte) {
	return byte(y >> (yBits - 8)), byte(y >> (yBits - 16) & 0xFF), byte(y >> (yBits - 24) & 0xFF), byte(y >> (yBits - 32) & 0xFF)
}

func (s PCM32bit) Encode(w io.Writer) {
	Encode(w, s, s.MaxX(), uint32(unitX/s.Period()), 4)
}

// Read a wave format stream into an array of PCMFunctions.
// one PCMFunction for each channel in the encoding.
func Decode(wav io.Reader) ([]PCMFunction, error) {
	bytesToRead, format, err := readHeader(wav)
	if err != nil {
		return nil, err
	}
	samples := bytesToRead / uint32(format.Channels) / uint32(format.Bits/8)
	sampleData, err := readData(wav, samples, uint32(format.Channels), uint32(format.Bits/8))
	if err != nil {
		return nil, err
	}
	functions := make([]PCMFunction, format.Channels)
	var c uint32
	for ; c < uint32(format.Channels); c++ {
		switch format.Bits {
		case 8:
			functions[c] = PCM8bit{PCM{unitX / x(format.SampleRate), unitX / x(format.SampleRate) * x(samples), sampleData[c*samples : (c+1)*samples]}}
		case 16:
			functions[c] = PCM16bit{PCM{unitX / x(format.SampleRate), unitX / x(format.SampleRate) * x(samples), sampleData[c*samples*2 : (c+1)*samples*2]}}
		case 24:
			functions[c] = PCM24bit{PCM{unitX / x(format.SampleRate), unitX / x(format.SampleRate) * x(samples), sampleData[c*samples*3 : (c+1)*samples*3]}}
		case 32:
			functions[c] = PCM32bit{PCM{unitX / x(format.SampleRate), unitX / x(format.SampleRate) * x(samples), sampleData[c*samples*4 : (c+1)*samples*4]}}
		}
	}
	return functions, nil
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
type formatChunk struct {
	Code        uint16
	Channels    uint16
	SampleRate  uint32
	ByteRate    uint32
	SampleBytes uint16
	Bits        uint16
}

func readHeader(wav io.Reader) (uint32, *formatChunk, error) {
	var header riffHeader
	var formatHeader chunkHeader
	var format formatChunk
	var dataHeader chunkHeader
	if err := binary.Read(wav, binary.LittleEndian, &header); err != nil {
		return 0, nil, ErrWavParse{"Header not complete."}
	}
	if header.C1 != 'R' || header.C2 != 'I' || header.C3 != 'F' || header.C4 != 'F' || header.C5 != 'W' || header.C6 != 'A' || header.C7 != 'V' || header.C8 != 'E' {
		return 0, nil, ErrWavParse{"Not RIFF/WAVE format."}
	}
	//var runningBytes int =16
	if err := binary.Read(wav, binary.LittleEndian, &formatHeader); err != nil {
		return 0, nil, ErrWavParse{"Chunk incomplete."}
	}
	// TODO skip other chunks
	if formatHeader.C1 != 'f' || formatHeader.C2 != 'm' || formatHeader.C3 != 't' || formatHeader.C4 != ' ' || formatHeader.DataLen != 16 {
		return 0, nil, ErrWavParse{"No format chunk."}
	}

	if err := binary.Read(wav, binary.LittleEndian, &format); err != nil {
		return 0, nil, ErrWavParse{"Format chunk incomplete."}
	}
	if format.Code != 1 {
		return 0, &format, errors.New("only PCM supported.")
	}
	if format.Channels == 0 || format.Channels > 2 {
		return 0, &format, errors.New("only mono or stereo PCM supported.")
	}
	if format.Bits%8 != 0 {
		return 0, &format, ErrWavParse{"not whole byte samples size!"}
	}

	//nice TODO a "LIST" chunk with, 3 fields third being "INFO", can contain "ICOP" and "ICRD" chunks providing copyright and creation date information.

	//	ByteRate    uint32
	//	SampleBytes uint16

	// skip any non-"data" chucks
	if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
		return 0, &format, ErrWavParse{"Chunk header incomplete."}
	}
	for dataHeader.C1 != 'd' || dataHeader.C2 != 'a' || dataHeader.C3 != 't' || dataHeader.C4 != 'a' {
		var err error
		if s, ok := wav.(io.Seeker); ok {
			_, err = s.Seek(int64(dataHeader.DataLen), os.SEEK_CUR) // seek relative to current file pointer
		} else {
			_, err = io.CopyN(ioutil.Discard, wav, int64(dataHeader.DataLen))
		}
		if err != nil {
			return 0, &format, ErrWavParse{string(dataHeader.C1) + string(dataHeader.C2) + string(dataHeader.C3) + string(dataHeader.C4) + " " + err.Error()}
		}

		if err := binary.Read(wav, binary.LittleEndian, &dataHeader); err != nil {
			return 0, &format, ErrWavParse{"Chunk header incomplete."}
		}
	}

	//if dataHeader.DataLen!=header.DataLen-36 {return nil, ErrWavParse{fmt.Sprintf("data chunk size mismatch. %v+36!=%v",dataHeader.DataLen,header.DataLen), []byte(fmt.Sprintf("%#v",dataHeader))}}	//  this is only true for non-extensible wav, ie non-microsoft
	if dataHeader.DataLen%uint32(format.Channels) != 0 {
		return 0, &format, ErrWavParse{fmt.Sprintf("sound sample data length %d not divisable by channel count", dataHeader.DataLen)}
	}
	return dataHeader.DataLen, &format, nil
}

func readData(wav io.Reader, samples uint32, channels uint32, sampleBytes uint32) ([]byte, error) {
	sampleData := make([]byte, samples*channels*sampleBytes)
	var s uint32
	var err error
	for ; s < samples; s++ {
		// deinterlace channels by reading directly into separate regions of a byte slice
		var c uint32
		for ; c < uint32(channels); c++ {
			if n, err := wav.Read(sampleData[(c*samples+s)*sampleBytes : (c*samples+s+1)*sampleBytes]); err != nil || n != int(sampleBytes) {
				return nil, ErrWavParse{fmt.Sprintf("data incomplete %v of %v", s, samples)}
			}
		}
	}
	return sampleData, err
}

