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
	buf := bufio.NewWriter(w)
	samplePeriod := X(1 / float32(sampleRate))
	samples := uint32(length/samplePeriod) + 1
	readerForPCM8Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM8bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset/samplePeriod) : uint32(offset.Offset/samplePeriod)+samples])
			} else if pcm, ok := s.(PCM8bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					_, err = w.Write([]byte{encodePCM8bit(s.property(x(i) * samplePeriod))})
				}
			}
			w.Close()
		}()
		return r
	}
	readerForPCM16Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM16bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset*2/samplePeriod) : uint32(offset.Offset*2/samplePeriod)+samples*2])
			} else if pcm, ok := s.(PCM16bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*2])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					b1, b2 := encodePCM16bit(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{b1, b2})
				}
			}
			w.Close()
		}()
		return r
	}
	readerForPCM24Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM24bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset*3/samplePeriod) : uint32(offset.Offset*3/samplePeriod)+samples*3])
			} else if pcm, ok := s.(PCM24bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*3])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					b1, b2, b3 := encodePCM24bit(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{ b1, b2,b3})
				}
			}
			w.Close()
		}()
		return r
	}
	readerForPCM32Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM32bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset*4/samplePeriod) : uint32(offset.Offset*4/samplePeriod)+samples*4])
			} else if pcm, ok := s.(PCM32bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*4])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					b1, b2, b3, b4 := encodePCM32bit(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{b1, b2, b3, b4})
				}
			}
			w.Close()
		}()
		return r
	}
	readerForPCM48Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM48bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset*6/samplePeriod) : uint32(offset.Offset*6/samplePeriod)+samples*6])
			} else if pcm, ok := s.(PCM48bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*6])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					b1, b2, b3, b4, b5, b6 := encodePCM48bit(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{ b1, b2, b3, b4,b5 ,b6})
				}
			}
			w.Close()
		}()
		return r
	}
	readerForPCM64Bit := func(s Signal) io.Reader {
		r, w := io.Pipe()
		go func() {
			// try shortcuts first
			offset, ok := s.(Offset)
			if pcms, ok2 := offset.LimitedSignal.(PCM64bit); ok && ok2 && pcms.samplePeriod == samplePeriod && pcms.MaxX() >= length-offset.Offset {
				w.Write(pcms.Data[uint32(offset.Offset*8/samplePeriod) : uint32(offset.Offset*8/samplePeriod)+samples*8])
			} else if pcm, ok := s.(PCM64bit); ok && pcm.samplePeriod == samplePeriod && pcm.MaxX() >= length {
				w.Write(pcm.Data[:samples*8])
			} else {
				defer func(){
					e:=recover()
					if e!=nil{
						w.CloseWithError(e.(error))
					}else{
						w.Close()
					}
				}()
				for i:=uint32(0); err ==nil && i < samples; i++ {
					b1, b2, b3, b4, b5, b6, b7, b8 := encodePCM64bit(s.property(x(i) * samplePeriod))
					_, err = w.Write([]byte{ b1, b2, b3, b4,b5 ,b6, b7, b8})
				}
			}
			w.Close()
		}()
		return r
	}
	binary.Write(buf, binary.LittleEndian, riffHeader{'R', 'I', 'F', 'F', samples*uint32(sampleBytes) + 36, 'W', 'A', 'V', 'E'})
	binary.Write(buf, binary.LittleEndian, chunkHeader{'f', 'm', 't', ' ', 16})
	binary.Write(buf, binary.LittleEndian, formatChunk{
		Code:        1,
		Channels:    uint16(len(ss)),
		SampleRate:  sampleRate,
		ByteRate:    sampleRate * uint32(sampleBytes) *uint32(len(ss)),
		SampleBytes: uint16(sampleBytes)*uint16(len(ss)),
		Bits:        uint16(8 * sampleBytes),
	})
	binary.Write(buf, binary.LittleEndian, chunkHeader{'d', 'a', 't', 'a', samples*uint32(sampleBytes)*uint32(len(ss))})
	readers:=make([]io.Reader,len(ss))
	switch sampleBytes {
	case 1:
		for i,_:=range(readers){
			readers[i]=readerForPCM8Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,1)
	case 2:
		for i,_:=range(readers){
			readers[i]=readerForPCM16Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,2)
	case 3:
		for i,_:=range(readers){
			readers[i]=readerForPCM24Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,3)
	case 4:
		for i,_:=range(readers){
			readers[i]=readerForPCM32Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,4)
	case 6:
		for i,_:=range(readers){
			readers[i]=readerForPCM48Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,6)
	case 8:
		for i,_:=range(readers){
			readers[i]=readerForPCM64Bit(ss[i])
		}
		err=interleavedWrite(buf,readers,8)
	}
	if err != nil {
		log.Println("Encode failure:" + err.Error())
	} else {
		buf.Flush()
	}
}


func interleavedWrite(w io.Writer, rs []io.Reader, blockSize int64) (err error) {
	if len(rs)==1{
		_,err=io.Copy(w, rs[0])
	}else if len(rs)>1{
		for err==nil{
			for i,_:=range(rs){
				_,err=io.CopyN(w,rs[i],blockSize)
			}
		}
		if err==io.EOF{err=nil}
	}
	return
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
	case PCM48bit:
		Encode(w, 6, uint32(unitX/s.Period()), p.MaxX(), p)
	case PCM64bit:
		Encode(w, 8, uint32(unitX/s.Period()), p.MaxX(), p)
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
	switch format.Bits {
	case 8:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM8bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples : (c+1)*samples]}}
		}
	case 16:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM16bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*2 : (c+1)*samples*2]}}
		}
	case 24:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM24bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*3 : (c+1)*samples*3]}}
		}
	case 32:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM32bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*4 : (c+1)*samples*4]}}
		}
	case 48:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM48bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*6 : (c+1)*samples*6]}}
		}
	case 64:
		for c:=uint32(0); c < uint32(format.Channels); c++ {
			pcms[c] = PCM64bit{PCM{unitX / x(format.SampleRate), sampleData[c*samples*8 : (c+1)*samples*8]}}
		}
	default:
			return nil,ErrWavParse{fmt.Sprintf("Unsupported bit depth (%d).",format.Bits)}
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
	var err error
	for s:=uint32(0); s < samples; s++ {
		// deinterlace channels by reading directly into separate regions of a byte slice
		for c:=uint32(0); c < uint32(channels); c++ {
			if n, err := wav.Read(sampleData[(c*samples+s)*sampleBytes : (c*samples+s+1)*sampleBytes]); err != nil || n != int(sampleBytes) {
				return nil, ErrWavParse{fmt.Sprintf("data incomplete %v of %v", s, samples)}
			}
		}
	}
	return sampleData, err
}



