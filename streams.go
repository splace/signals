package signals

import (
	"errors"
	"io"
	"net/http"
	"encoding/gob"
)
func init() {
	gob.Register(&Wave{})
}

const bufferSize = 16 

// a PCM-Signal read, as required, from a URL.
// if queried for a property value for an x that is more than 32 samples lower than a previous query, will return zero.
type Wave struct{
	Shifted
	URL string
	reader io.Reader
}

func NewWave(URL string) (*Wave, error) {
	r, channels, bytes, rate, err := PCMReader(URL)
	if err!=nil {
		return nil,err
	}	
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
	case 8:
		return &Wave{Shifted{NewPCM64bit(rate, b),0},URL,r}, nil
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
	for offset > s.MaxX() {
		// append available data onto the PCM slice.
		// also possibly shift off some data, shortening the PCM slice, retaining at least two buffer lengths.
		// semi-samples are read but not accessed by property.
		switch st:=s.Shifted.Signal.(type) {
		case PCM8bit:
			sd:=PCM8bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize+n]
			if len(sd.Data)>bufferSize*3{
				sd.Data=sd.Data[bufferSize:]
				s.Shifted.Shift+=bufferSize*st.samplePeriod
			}
			s.Shifted=Shifted{sd,s.Shifted.Shift}
		case PCM16bit:
			sd:=PCM16bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize*2)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*2:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize*2+n]
			if len(sd.Data)>bufferSize*2*3{
				sd.Data=sd.Data[bufferSize*2:]
				s.Shifted=Shifted{sd,s.Shifted.Shift+bufferSize*st.samplePeriod}
				}else{
				s.Shifted=Shifted{sd,s.Shifted.Shift}
			}
		case PCM24bit:
			sd:=PCM24bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize*3)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*3:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize*3+n]
			if len(sd.Data)>bufferSize*3*3{
				sd.Data=sd.Data[bufferSize*3:]
				s.Shifted=Shifted{sd,s.Shifted.Shift+bufferSize*st.samplePeriod}
				}else{
				s.Shifted=Shifted{sd,s.Shifted.Shift}
			}
		case PCM32bit:
			sd:=PCM16bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize*4)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*4:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize*4+n]
			if len(sd.Data)>bufferSize*4*3{
				sd.Data=sd.Data[bufferSize*4:]
				s.Shifted=Shifted{sd,s.Shifted.Shift+bufferSize*st.samplePeriod}
				}else{
				s.Shifted=Shifted{sd,s.Shifted.Shift}
			}
		case PCM48bit:
			sd:=PCM48bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize*6)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*6:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize*6+n]
			if len(sd.Data)>bufferSize*6*3{
				sd.Data=sd.Data[bufferSize*6:]
				s.Shifted=Shifted{sd,s.Shifted.Shift+bufferSize*st.samplePeriod}
				}else{
				s.Shifted=Shifted{sd,s.Shifted.Shift}
			}
		case PCM64bit:
			sd:=PCM64bit{st.PCM}
			sd.Data=append(sd.Data,make([]byte,bufferSize*8)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*8:])
			failOn(err)
			sd.Data=sd.Data[:len(sd.Data)-bufferSize*8+n]
			if len(sd.Data)>bufferSize*8*3{
				sd.Data=sd.Data[bufferSize*8:]
				s.Shifted=Shifted{sd,s.Shifted.Shift+bufferSize*st.samplePeriod}
				}else{
				s.Shifted=Shifted{sd,s.Shifted.Shift}
			}
		}
	}
	return s.Shifted.property(offset)
}


//func updateShifted(s Shifted, r io.Reader, b *[]byte, blockSize int) (err error){
//	b=append(b,make([]byte,bufferSize*blockSize)...)
//	n, err := r.Read(b[len(b)-bufferSize*blockSize:])
//	failOn(err)
//	b=b[:len(b)-bufferSize*blockSize+n]
//	if len(b)>bufferSize*blockSize*3{
//		b=b[bufferSize*blockSize:]
//		s.Shift+=bufferSize*s.samplePeriod
//	}
//}

func PCMReader(source string) (io.Reader, uint16, uint16, uint32, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, 0, 0, 0, err
	}
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
