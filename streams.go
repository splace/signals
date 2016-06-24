package signals

import (
	"errors"
	"io"
	"net/http"
	"encoding/gob"
	//"fmt"
)
func init() {
	gob.Register(&Wave{})
}

const bufferSize = 16 

// a PCM-Signal read, as required, from a URL.
// if queried for its property value for an x that is more than 32 samples lower than a previous query, will return zero.
type Wave struct{
	Shifted
	URL string
	reader io.Reader
}

func NewWave(URL string) (*Wave, error) {
	r, channels, bytes, rate, err := PCMReader(URL)
	failOn(err)
	if channels != 1 {
		return nil, errors.New(URL+":Needs to be mono.")
	}
	b := make([]byte, bufferSize*bytes)
	n, err := r.Read(b)
	failOn(err)
	b=b[:n]
	switch bytes {
	case 1:
		return &Wave{Shifted{NewPCM8bit(rate, b),0},URL,r}, nil
	case 2:
		return &Wave{Shifted{NewPCM16bit(rate, b),0},URL,r}, nil
	case 3:
		return &Wave{Shifted{NewPCM24bit(rate, b),0},URL,r}, nil
	case 4:
		return &Wave{Shifted{NewPCM32bit(rate, b),0},URL,r}, nil
	case 6:
		return &Wave{Shifted{NewPCM48bit(rate, b),0},URL,r}, nil
	}
	return nil, ErrWavParse{"Source bit rate not supported."}
}

func (s *Wave) property(offset x) y {
	if s.reader==nil{
		wav,err:=NewWave(s.URL)
		failOn(err)
		s.Shifted=wav.Shifted
		s.reader=wav.reader
	}
	if offset > s.MaxX() {
		switch st:=s.Shifted.Signal.(type) {
		case PCM8bit:
			b:=make([]byte,bufferSize)
			st.Data=append(st.Data,b...)
			n, err := s.reader.Read(st.Data[len(st.Data)-bufferSize:])
			failOn(err)
			st.Data=st.Data[:len(st.Data)-bufferSize+n]
			if len(st.Data)>bufferSize*3{
				st.Data=st.Data[bufferSize:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
		case PCM16bit:
			b:=make([]byte,bufferSize*2)
			st.Data=append(st.Data,b...)
			n, err := s.reader.Read(st.Data[len(st.Data)-bufferSize*2:])
			failOn(err)
			st.Data=st.Data[:len(st.Data)-bufferSize*2+n]
			if len(st.Data)>bufferSize*6{
				st.Data=st.Data[bufferSize*2:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
		case PCM24bit:
			b:=make([]byte,bufferSize*3)
			st.Data=append(st.Data,b...)
			n, err := s.reader.Read(st.Data[len(st.Data)-bufferSize*3:])
			failOn(err)
			st.Data=st.Data[:len(st.Data)-bufferSize*3+n]
			if len(st.Data)>bufferSize*9{
				st.Data=st.Data[bufferSize*3:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
		case PCM32bit:
			b:=make([]byte,bufferSize*4)
			st.Data=append(st.Data,b...)
			n, err := s.reader.Read(st.Data[len(st.Data)-bufferSize*4:])
			failOn(err)
			st.Data=st.Data[:len(st.Data)-bufferSize*4+n]
			if len(st.Data)>bufferSize*12{
				st.Data=st.Data[bufferSize*4:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
		case PCM48bit:
			b:=make([]byte,bufferSize*6)
			st.Data=append(st.Data,b...)
			n, err := s.reader.Read(st.Data[len(st.Data)-bufferSize*6:])
			failOn(err)
			st.Data=st.Data[:len(st.Data)-bufferSize*6+n]
			if len(st.Data)>bufferSize*18{
				st.Data=st.Data[bufferSize*6:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
		}
	}
	return s.Shifted.property(offset)
}

func PCMReader(source string) (io.Reader, uint16, uint16, uint32, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	//fmt.Println(resp.Header)
	if resp.Header["Content-Type"][0] == "sound/wav" || resp.Header["Content-Type"][0] == "audio/x-wav" {
		_, format, err := readHeader(resp.Body)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		return resp.Body, format.Channels, format.SampleBytes, format.SampleRate, nil
	}
	if resp.Header["Content-Type"][0] == "audio/l16;rate=8000" {
		return resp.Body, 1, 2, 8000, nil
	}
	return nil, 0, 0, 0, errors.New("Source in unrecognized format.")
}

func failOn(e error){
	if e!=nil {panic(e)}
}
