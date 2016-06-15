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

type formatChunk struct {
	Code        uint16
	Channels    uint16
	SampleRate  uint32
	ByteRate    uint32
	SampleBytes uint16
	Bits        uint16
}

// Encode Signals as PCM data,in a Riff wave container.
func Encode(w io.Writer, sampleBytes uint8, sampleRate uint32, length x, ss ...Signal) {
	var err error
	var i uint32
	buf := bufio.NewWriter(w)
	samplePeriod := X(1 / float32(sampleRate))
	samples := uint32(length/samplePeriod) + 1
	readPCM8Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			shifted, ok := s.(Shifted)
			if pcms, ok2 := shifted.Signal.(PCM8bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-shifted.Shift {
				w.Write(pcms.Data[uint32(shifted.Shift/samplePeriod) : uint32(shifted.Shift/samplePeriod)+samples])
			} else if pcm, ok := s.(PCM8bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples])
			} else {
				for ; i < samples; i++ {
					_, err = w.Write([]byte{PCM8bitEncode(s.property(x(i) * samplePeriod))})
					if err != nil {
						break
					}
				}
			}
			w.Close()
		}()
		return r
	}
	readPCM16Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			shifted, ok := s.(Shifted)
			if pcms, ok2 := shifted.Signal.(PCM16bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-shifted.Shift {
				w.Write(pcms.Data[uint32(shifted.Shift*2/samplePeriod) : uint32(shifted.Shift*2/samplePeriod)+samples*2])
			} else if pcm, ok := s.(PCM16bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*2])
			} else {
				for ; i < samples; i++ {
					b1, b2 := PCM16bitEncode(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{b2, b1})
					if err != nil {
						break
					}
				}
			}
			w.Close()
		}()
		return r
	}
	readPCM24Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			shifted, ok := s.(Shifted)
			if pcms, ok2 := shifted.Signal.(PCM24bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-shifted.Shift {
				w.Write(pcms.Data[uint32(shifted.Shift*3/samplePeriod) : uint32(shifted.Shift*3/samplePeriod)+samples*3])
			} else if pcm, ok := s.(PCM24bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*3])
			} else {
				for ; i < samples; i++ {
					b1, b2, b3 := PCM24bitEncode(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{b3, b2, b1})
					if err != nil {
						break
					}
				}
			}
			w.Close()
		}()
		return r
	}
	readPCM32Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			shifted, ok := s.(Shifted)
			if pcms, ok2 := shifted.Signal.(PCM32bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-shifted.Shift {
				w.Write(pcms.Data[uint32(shifted.Shift*4/samplePeriod) : uint32(shifted.Shift*4/samplePeriod)+samples*4])
			} else if pcm, ok := s.(PCM32bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*4])
			} else {
				for ; i < samples; i++ {
					b1, b2, b3, b4 := PCM32bitEncode(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{b4, b3, b2, b1})
					if err != nil {
						break
					}
				}
			}
			w.Close()
		}()
		return r
	}
	binary.Write(w, binary.LittleEndian, riffHeader{'R', 'I', 'F', 'F', samples*uint32(sampleBytes) + 36, 'W', 'A', 'V', 'E'})
	binary.Write(w, binary.LittleEndian, chunkHeader{'f', 'm', 't', ' ', 16})
	binary.Write(w, binary.LittleEndian, formatChunk{
		Code:        1,
		Channels:    uint16(len(ss)),
		SampleRate:  sampleRate,
		ByteRate:    sampleRate * uint32(sampleBytes) *uint32(len(ss)),
		SampleBytes: uint16(sampleBytes)*uint16(len(ss)),
		Bits:        uint16(8 * sampleBytes)*uint16(len(ss)),
	})
	fmt.Fprint(w, "data")
	binary.Write(w, binary.LittleEndian, samples*uint32(sampleBytes))
	readers:=make([]io.Reader,len(ss))
	switch sampleBytes {
	case 1:
		for i,_:=range(readers){
			readers[i]=readPCM8Bit(ss[i])
		}
		if len(readers)==1{
			_,err=io.Copy(buf, readers[0])
			if err!=nil{
				panic(err)
			}
		}else{
			for err==nil{
				for i,_:=range(readers){
					_,err=io.CopyN(buf,readers[i],1)
				}
			}
			if err==io.EOF{err=nil}
		}
	case 2:
		for i,_:=range(readers){
			readers[i]=readPCM16Bit(ss[i])
		}
		if len(readers)==1{
			_,err=io.Copy(buf, readers[0])
			if err!=nil{
				panic(err)
			}
		}else{
			for err==nil{
				for i,_:=range(readers){
					_,err=io.CopyN(buf,readers[i],2)
				}
			}
			if err==io.EOF{err=nil}
		}
	case 3:
		for i,_:=range(readers){
			readers[i]=readPCM24Bit(ss[i])
		}
		if len(readers)==1{
			_,err=io.Copy(buf, readers[0])
			if err!=nil{
				panic(err)
			}
		}else{
			for err==nil{
				for i,_:=range(readers){
					_,err=io.CopyN(buf,readers[i],3)
				}
			}
			if err==io.EOF{err=nil}
		}
	case 4:
		for i,_:=range(readers){
			readers[i]=readPCM32Bit(ss[i])
		}
		if len(readers)==1{
			_,err=io.Copy(buf, readers[0])
			if err!=nil{
				panic(err)
			}
		}else{
			for err==nil{
				for i,_:=range(readers){
					_,err=io.CopyN(buf,readers[i],4)
				}
			}
			if err==io.EOF{err=nil}
		}
	}
	if err != nil {
		log.Println("Encode failure:" + err.Error() + fmt.Sprint(w))
	} else {
		buf.Flush()
	}
}

// encode a LimitedSignal with a sampleRate equal to the Period() of a given PeriodicSignal, and its precision if its a PCM type, otherwise defaults to 16bit.
func EncodeLike(w io.Writer, s PeriodicSignal, p LimitedSignal) {
	switch s.(type) {
	case PCM8bit:
		Encode(w, 1, uint32(unitX/s.Period()), p.MaxX(), p)
	case PCM16bit:
		Encode(w, 2, uint32(unitX/s.Period()), p.MaxX(), p)
	case PCM24bit:
		Encode(w, 3, uint32(unitX/s.Period()), p.MaxX(), p)
	case PCM32bit:
		Encode(w, 4, uint32(unitX/s.Period()), p.MaxX(), p)
	default:
		Encode(w, 2, uint32(unitX/s.Period()), p.MaxX(), p)
	}
	return
}

// Read a wave format stream into an array of PeriodicLimitedSignals.
// one for each channel in the encoding.
func Decode(wav io.Reader) ([]PeriodicLimitedSignal, error) {
	bytesToRead, format, err := readHeader(wav)
	if err != nil {
		return nil, err
	}
	samples := bytesToRead / uint32(format.Channels) / uint32(format.Bits/8)
	sampleData, err := readData(wav, samples, uint32(format.Channels), uint32(format.Bits/8))
	if err != nil {
		return nil, err
	}
	pcms := make([]PeriodicLimitedSignal, format.Channels)
	var c uint32
	for ; c < uint32(format.Channels); c++ {
		switch format.Bits {
		case 8:
			pcms[c] = PCM8bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples : (c+1)*samples]}}
		case 16:
			pcms[c] = PCM16bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*2 : (c+1)*samples*2]}}
		case 24:
			pcms[c] = PCM24bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*3 : (c+1)*samples*3]}}
		case 32:
			pcms[c] = PCM32bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*4 : (c+1)*samples*4]}}
		}
	}
	return pcms, nil
}

type ErrWavParse struct {
	description string
}

func (e ErrWavParse) Error() string {
	return fmt.Sprintf("WAVE Parse,%s", e.description)
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
	//if format.Channels == 0 || format.Channels > 2 {
	//	return 0, &format, errors.New("only mono or stereo PCM supported.")
	//}
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



/*  Hal3 Wed Jun 15 00:02:29 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestNoiseSave
--- PASS: TestNoiseSave (0.92s)
=== RUN   TestSaveWav
--- PASS: TestSaveWav (0.03s)
=== RUN   TestLoad
--- PASS: TestLoad (0.02s)
=== RUN   TestLoadChannels
--- PASS: TestLoadChannels (0.08s)
=== RUN   TestMultiChannelSave
--- PASS: TestMultiChannelSave (0.75s)
=== RUN   TestStackPCMs
--- PASS: TestStackPCMs (0.33s)
=== RUN   TestMultiplexTones
--- PASS: TestMultiplexTones (0.12s)
=== RUN   TestSaveLoadSave
--- PASS: TestSaveLoadSave (0.15s)
=== RUN   TestPiping
--- PASS: TestPiping (0.02s)
=== RUN   TestRawPCM
--- PASS: TestRawPCM (0.00s)
=== RUN   TestSplitPCM
--- PASS: TestSplitPCM (0.00s)
=== RUN   TestEnocdePCMToShortLength
--- PASS: TestEnocdePCMToShortLength (0.00s)
=== RUN   TestEnocdeShiftedPCM
--- PASS: TestEnocdeShiftedPCM (0.00s)
=== RUN   TestImagingSine
--- PASS: TestImagingSine (0.27s)
=== RUN   TestImaging
--- PASS: TestImaging (0.28s)
=== RUN   TestComposable
--- PASS: TestComposable (1.50s)
=== RUN   TestStackimage
--- PASS: TestStackimage (0.93s)
=== RUN   TestMultiplexImage
--- PASS: TestMultiplexImage (0.91s)
=== RUN   ExampleADSREnvelope
--- PASS: ExampleADSREnvelope (0.00s)
=== RUN   ExamplePulsePattern
--- PASS: ExamplePulsePattern (0.00s)
=== RUN   ExampleNoise
--- PASS: ExampleNoise (0.00s)
=== RUN   ExampleConstantZero
--- PASS: ExampleConstantZero (0.00s)
=== RUN   ExampleConstantUnity
--- PASS: ExampleConstantUnity (0.00s)
=== RUN   ExampleSquare
--- PASS: ExampleSquare (0.00s)
=== RUN   ExamplePulse
--- PASS: ExamplePulse (0.00s)
=== RUN   ExampleRampUpDown
--- PASS: ExampleRampUpDown (0.00s)
=== RUN   ExampleHeavyside
--- PASS: ExampleHeavyside (0.00s)
=== RUN   ExampleSine
--- PASS: ExampleSine (0.00s)
=== RUN   ExampleSigmoid
--- PASS: ExampleSigmoid (0.00s)
=== RUN   ExampleShifted
--- PASS: ExampleShifted (0.00s)
=== RUN   ExampleReflected
--- PASS: ExampleReflected (0.00s)
=== RUN   ExamplePower
--- PASS: ExamplePower (0.00s)
=== RUN   ExampleModulated
--- PASS: ExampleModulated (0.00s)
=== RUN   ExampleStack
--- PASS: ExampleStack (0.00s)
=== RUN   ExampleTriggered
--- PASS: ExampleTriggered (0.00s)
=== RUN   ExampleSegmented
--- PASS: ExampleSegmented (0.00s)
=== RUN   ExampleSegmented_makeSawtooth
--- PASS: ExampleSegmented_makeSawtooth (0.00s)
=== RUN   ExampleRateModulated
--- PASS: ExampleRateModulated (0.00s)
=== RUN   ExampleLooped
--- PASS: ExampleLooped (0.00s)
=== RUN   ExampleRepeated
--- PASS: ExampleRepeated (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	6.381s
Wed Jun 15 00:02:38 BST 2016 */
/*  Hal3 Wed Jun 15 00:06:23 BST 2016 go version go1.5.1 linux/amd64
FAIL	_/home/simon/Dropbox/github/working/signals [build failed]
Wed Jun 15 00:06:24 BST 2016 */
/*  Hal3 Wed Jun 15 00:06:32 BST 2016 go version go1.5.1 linux/amd64
=== RUN   TestNoiseSave
--- PASS: TestNoiseSave (0.91s)
=== RUN   TestSaveWav
--- PASS: TestSaveWav (0.03s)
=== RUN   TestLoad
--- PASS: TestLoad (0.02s)
=== RUN   TestLoadChannels
--- PASS: TestLoadChannels (0.07s)
=== RUN   TestMultiChannelSave
--- PASS: TestMultiChannelSave (0.76s)
=== RUN   TestStackPCMs
--- PASS: TestStackPCMs (0.40s)
=== RUN   TestMultiplexTones
--- PASS: TestMultiplexTones (0.11s)
=== RUN   TestSaveLoadSave
--- PASS: TestSaveLoadSave (0.14s)
=== RUN   TestPiping
--- PASS: TestPiping (0.02s)
=== RUN   TestRawPCM
--- PASS: TestRawPCM (0.00s)
=== RUN   TestSplitPCM
--- PASS: TestSplitPCM (0.00s)
=== RUN   TestEnocdePCMToShortLength
--- PASS: TestEnocdePCMToShortLength (0.00s)
=== RUN   TestEnocdeShiftedPCM
--- PASS: TestEnocdeShiftedPCM (0.00s)
=== RUN   TestImagingSine
--- PASS: TestImagingSine (0.30s)
=== RUN   TestImaging
--- PASS: TestImaging (0.33s)
=== RUN   TestComposable
--- PASS: TestComposable (1.57s)
=== RUN   TestStackimage
--- PASS: TestStackimage (0.92s)
=== RUN   TestMultiplexImage
--- PASS: TestMultiplexImage (0.91s)
=== RUN   ExampleADSREnvelope
--- PASS: ExampleADSREnvelope (0.00s)
=== RUN   ExamplePulsePattern
--- PASS: ExamplePulsePattern (0.00s)
=== RUN   ExampleNoise
--- PASS: ExampleNoise (0.01s)
=== RUN   ExampleConstantZero
--- PASS: ExampleConstantZero (0.00s)
=== RUN   ExampleConstantUnity
--- PASS: ExampleConstantUnity (0.00s)
=== RUN   ExampleSquare
--- PASS: ExampleSquare (0.00s)
=== RUN   ExamplePulse
--- PASS: ExamplePulse (0.00s)
=== RUN   ExampleRampUpDown
--- PASS: ExampleRampUpDown (0.00s)
=== RUN   ExampleHeavyside
--- PASS: ExampleHeavyside (0.00s)
=== RUN   ExampleSine
--- PASS: ExampleSine (0.00s)
=== RUN   ExampleSigmoid
--- PASS: ExampleSigmoid (0.00s)
=== RUN   ExampleShifted
--- PASS: ExampleShifted (0.00s)
=== RUN   ExampleReflected
--- PASS: ExampleReflected (0.00s)
=== RUN   ExamplePower
--- PASS: ExamplePower (0.00s)
=== RUN   ExampleModulated
--- PASS: ExampleModulated (0.00s)
=== RUN   ExampleStack
--- PASS: ExampleStack (0.00s)
=== RUN   ExampleTriggered
--- PASS: ExampleTriggered (0.00s)
=== RUN   ExampleSegmented
--- PASS: ExampleSegmented (0.00s)
=== RUN   ExampleSegmented_makeSawtooth
--- PASS: ExampleSegmented_makeSawtooth (0.00s)
=== RUN   ExampleRateModulated
--- PASS: ExampleRateModulated (0.00s)
=== RUN   ExampleLooped
--- PASS: ExampleLooped (0.00s)
=== RUN   ExampleRepeated
--- PASS: ExampleRepeated (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	6.544s
Wed Jun 15 00:06:41 BST 2016 */

