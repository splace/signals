package signals

import (
	"encoding/base64"
	"encoding/gob"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	gob.Register(&Wave{})
}

const bufferSize = 16

// an offset PCM Signal, (so single channel) that reads from a source, as required, its data.
// supported URL schemes, "file:", "data:", "http(s):".
// Http(s); MIME: "audio/l?;rate=?","sound/wav"(mono),"audio/x-wav" (mono)
// File: ".wav"(mono),".pcm",".gob"
// Data: MIME: "audio/l?;rate=?","sound/wav"(mono),"audio/x-wav" (mono), ENCODING: "base64" or none.
// if queried for a property value from an x that is more than 32 samples lower than a previous query, might return zero.
type Wave struct {
	Offset
	URL    string
	reader io.Reader
}

func (s *Wave) property(p x) y {
	if s.reader == nil {
		wav, err := NewWave(s.URL)
		failOn(err)
		s.Offset = wav.Offset
		s.reader = wav.reader
	}
	for p > s.MaxX() {
		// append available data onto the PCM slice.
		// also possibly shift off some data, shortening the PCM slice, retaining at least two buffer lengths.
		// partial samples are read but not accessed by property.
		switch st := s.Offset.LimitedSignal.(type) {
		case PCM8bit:
			sd := PCM8bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize+n]
			if len(sd.Data) > bufferSize*3 {
				sd.Data = sd.Data[bufferSize:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		case PCM16bit:
			sd := PCM16bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize*2)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*2:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize*2+n]
			if len(sd.Data) > bufferSize*2*3 {
				sd.Data = sd.Data[bufferSize*2:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		case PCM24bit:
			sd := PCM24bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize*3)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*3:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize*3+n]
			if len(sd.Data) > bufferSize*3*3 {
				sd.Data = sd.Data[bufferSize*3:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		case PCM32bit:
			sd := PCM32bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize*4)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*4:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize*4+n]
			if len(sd.Data) > bufferSize*4*3 {
				sd.Data = sd.Data[bufferSize*4:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		case PCM48bit:
			sd := PCM48bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize*6)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*6:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize*6+n]
			if len(sd.Data) > bufferSize*6*3 {
				sd.Data = sd.Data[bufferSize*6:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		case PCM64bit:
			sd := PCM64bit{st.PCM}
			sd.Data = append(sd.Data, make([]byte, bufferSize*8)...)
			n, err := s.reader.Read(sd.Data[len(sd.Data)-bufferSize*8:])
			failOn(err)
			sd.Data = sd.Data[:len(sd.Data)-bufferSize*8+n]
			if len(sd.Data) > bufferSize*8*3 {
				sd.Data = sd.Data[bufferSize*8:]
				s.Offset = Offset{sd, s.Offset.Offset + bufferSize*st.samplePeriod}
			} else {
				s.Offset = Offset{sd, s.Offset.Offset}
			}
		}
	}
	return s.Offset.property(p)
}

//func updateShifted(s Shifted, r io.Reader, b *[]byte, blockSize int) (err error){
//	b=append(b,make([]byte,bufferSize*blockSize)...)
//	n, err := r.Read(b[len(b)-bufferSize*blockSize:])
//	failOn(err)
//	b=b[:len(b)-bufferSize*blockSize+n]
//	if len(b)>bufferSize*blockSize*3{
//		b=b[bufferSize*blockSize:]
//		s.Offset+=bufferSize*s.samplePeriod
//	}
//}

func NewWave(URL string) (*Wave, error) {
	r, channels, bytes, rate, err := pcmReader(URL)
	if err != nil {
		return nil, err
	}
	if channels != 1 {
		return nil, errors.New(URL + ":Needs to be mono.")
	}
	b := make([]byte, bufferSize*bytes)
	n, err := r.Read(b)
	failOn(err)
	b = b[:n]
	switch bytes {
	case 1:
		return &Wave{Offset{NewPCM8bit(rate, b), 0}, URL, r}, nil
	case 2:
		return &Wave{Offset{NewPCM16bit(rate, b), 0}, URL, r}, nil
	case 3:
		return &Wave{Offset{NewPCM24bit(rate, b), 0}, URL, r}, nil
	case 4:
		return &Wave{Offset{NewPCM32bit(rate, b), 0}, URL, r}, nil
	case 6:
		return &Wave{Offset{NewPCM48bit(rate, b), 0}, URL, r}, nil
	case 8:
		return &Wave{Offset{NewPCM64bit(rate, b), 0}, URL, r}, nil
	}
	return nil, errors.New("Sample Bytes not supported:"+URL)
}

var contentTypeParse = regexp.MustCompile(`^audio/l(\d+);rate=(\d+)$`)

// returns a reader to a resource, along with its Channel count, Precision (bytes) and Sample rate.
func pcmReader(resourceLocation string) (io.Reader, uint16, uint16, uint32, error) {
	//	resp, err := http.Get(resourceLocation)
	url, err := url.Parse(resourceLocation)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	switch url.Scheme {
	case "file":
		switch path.Ext(url.Path) {
		case ".wav", ".wave", ".WAV", ".WAVE":
			file, err := os.Open(url.Path)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			_, format, err := readWaveHeader(file)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return file, format.Channels, format.SampleBytes, format.SampleRate, nil
		case ".gob", ".GOB":
			s, err := LoadGOB(url.Path[:len(url.Path)-4])
			if err != nil {
				return nil, 0, 0, 0, err
			}
			var sampleRate uint32 = 22010
			samplePeriod := X(1 / float32(sampleRate))
			r, w := io.Pipe()
			go func() {
				defer func() {
					e := recover()
					if e != nil {
						w.CloseWithError(e.(error))
					} else {
						w.Close()
					}
				}()
				for i, sample := uint32(0), make([]byte, 2); err == nil; i++ {
					sample[0], sample[1] = encodePCM16bit(s.property(x(i) * samplePeriod))
					_, err = w.Write(sample)
				}
			}()
			return r, 1, 2, sampleRate, nil
		case ".pcm":
			rate, err := strconv.ParseUint(path.Base(path.Dir(url.Path)), 10, 32)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			bits, err := strconv.ParseUint(path.Base(path.Dir(path.Dir(url.Path))[:len(path.Dir(path.Dir(url.Path)))-3]), 10, 20)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			file, err := os.Open(url.Path)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return file, 1, uint16(bits / 8), uint32(rate), nil
		}
	case "data":
		mimeAndRest := strings.SplitN(url.Opaque, ";", 2)
		encodingAndData := strings.SplitN(mimeAndRest[1], ",", 2)
		var r io.Reader
		if encodingAndData[0] == "base64" {
			r = base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodingAndData[1]))
		} else {
			r = strings.NewReader(encodingAndData[1])
		}
		if mimeAndRest[0] == "sound/wav" || mimeAndRest[0] == "audio/x-wav" {
			_, format, err := readWaveHeader(r)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return r, format.Channels, format.SampleBytes, format.SampleRate, nil
		}
		pcmFormat := contentTypeParse.FindStringSubmatch(mimeAndRest[0])
		if pcmFormat != nil {
			bits, err := strconv.ParseUint(pcmFormat[1], 10, 19)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			rate, err := strconv.ParseUint(pcmFormat[2], 10, 32)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return r, 1, uint16(bits / 8), uint32(rate), nil
		}
	default: // whatever supported and placed in Body, currently basically "http" or "https"
		resp, err := http.DefaultClient.Do(&http.Request{Method: "GET", URL: url})

		if err != nil {
			return nil, 0, 0, 0, err
		}
		if resp.Header["Content-Type"][0] == "sound/wav" || resp.Header["Content-Type"][0] == "audio/x-wav" {
			_, format, err := readWaveHeader(resp.Body)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return resp.Body, format.Channels, format.SampleBytes, format.SampleRate, nil
		}
		pcmFormat := contentTypeParse.FindStringSubmatch(resp.Header["Content-Type"][0])
		if pcmFormat != nil {
			bits, err := strconv.ParseUint(pcmFormat[1], 10, 19)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			rate, err := strconv.ParseUint(pcmFormat[2], 10, 32)
			if err != nil {
				return nil, 0, 0, 0, err
			}
			return resp.Body, 1, uint16(bits / 8), uint32(rate), nil
		}
	}
	return nil, 0, 0, 0, errors.New("Source:" + resourceLocation + " unsupported.")
}

func failOn(e error) {
	if e != nil {
		panic(e)
	}
}


