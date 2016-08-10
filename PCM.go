package signals

import (
	"io"
	"os"
	"io/ioutil"
	"path"
	"strconv"
	"errors"
)

// PCM is the state and behaviour common to all PCM, it doesn't include encoding information, so cannot return a property and so is not a Signal.
// Specific PCM<<encoding>> types embed this, and then are Signal's.
// these specific precision types, the Signals, return continuous property values that step from one PCM value to the next, Segmented could be used to get interpolated property values.
type PCM struct {
	samplePeriod x
	Data         []byte
}

// make a PCM type, from raw bytes.
func NewPCM(sampleRate uint32, Data []byte) PCM {
	return PCM{X(1 / float32(sampleRate)), Data}
}

func (s PCM) Period() x {
	return s.samplePeriod
}

//func ReadPCM(pathTo string) (p *PCM,err error) {)

// load a PCM from <<pathTo>>, which if not explicitly pointing into a folder with the <<Sample Rate>>, numerically as its name, will look into a sub-folder with the sampleRate indicated by the PCM parameter, and if that's zero will load any samplerate available. Also adds extension ".pcm".
func LoadPCM(pathTo string, p *PCM) (err error) {
	sampleRate,err:=strconv.ParseUint(path.Base(path.Dir(pathTo)), 10, 32)
	if err!=nil {
		if p.samplePeriod==0 {
			var files []os.FileInfo
			files, err = ioutil.ReadDir(path.Dir(pathTo))
			if err != nil {return}
			for _, file := range files {
				if file.IsDir() && file.Size()>0 {
					sampleRate,err=strconv.ParseUint(file.Name(), 10, 32)
					if err==nil {
						pathTo=path.Join(path.Dir(pathTo),file.Name(),path.Base(pathTo))
						break
					}
				}
			}
		}else{
			pathTo=path.Join(path.Dir(pathTo),strconv.FormatInt(int64(unitX / x(p.samplePeriod)),10),path.Base(pathTo))
		}
	}else{
		p.samplePeriod=X(1 / float32(sampleRate))
	}	
	p.Data,err=ioutil.ReadFile(pathTo+".pcm")
	return 
}

// save PCM into a paths subfolder depending on its sample rate. (see LoadPCM) 
func SavePCM(path string,pcm PCM) error {
	return pcm.SaveTo(path)
}

// save a PCM to <<pathTo>>, which if not inside a folder with the PCM's Sample Rate as its name, will add it, (making a new folder if required) which means the file won't then actually simply be at the <<pathTo>> address, but the LoadPCM function can automatically find the sub-folder. Also adds extension ".pcm".
func (p PCM) SaveTo(pathTo string) error {
	sampleRate,err:=strconv.ParseUint(path.Base(path.Dir(pathTo)), 10, 32)
	if err!=nil {
		pathTo=path.Join(path.Dir(pathTo),strconv.FormatInt(int64(unitX / x(p.samplePeriod)),10),path.Base(pathTo))
		err:=os.Mkdir(path.Dir(pathTo), os.ModeDir | 0775)
		if err!=nil && !os.IsExist(err){return err}
	}else{
		if p.samplePeriod!=unitX / x(sampleRate) {return errors.New("parent folder for different sample rate.")}
	}	
	file, err := os.Create(pathTo+".pcm")
	if err!=nil {return err}
	file.Write(p.Data)
	return file.Close()
}



// from a PCM return two new PCM's (with the same underlying data) from either side of a sample.
func (s PCM) Split(sample uint32, sampleBytes uint8) (head PCM, tail PCM) {
	copy := func(s PCM) PCM { return s }
	bytePosition := sample * uint32(sampleBytes)
	if bytePosition > uint32(len(s.Data)) {
		bytePosition = uint32(len(s.Data))
	}
	head, tail = s, copy(s)
	tail.Data = tail.Data[bytePosition:]
	head.Data = head.Data[:bytePosition]
	return
}

// 8 bit PCM Signal.
type PCM8bit struct {
	PCM
}

func NewPCM8bit(sampleRate uint32, Data []byte) PCM8bit {
	return PCM8bit{NewPCM(sampleRate, Data)}
}

func (s PCM8bit) property(p x) y {
	index := int(p / s.samplePeriod)
	if index < 0 || index >= len(s.Data){
		return 0
	}
	return decodePCM8bit(s.Data[index])
}


func encodePCM8bit(v y) byte {
	return byte(v>>(yBits-8)) + 128
}

func decodePCM8bit(b byte) y {
	return y(b-128) << (yBits-8)
}

func (s PCM8bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-1)
}

func (s PCM8bit) Encode(w io.Writer) {
	Encode(w, 1, uint32(unitX/s.Period()), s.MaxX(), s)
}

func (s PCM8bit) Split(p x) (PCM8bit, PCM8bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 1)
	return PCM8bit{head}, PCM8bit{tail}
}

// 16 bit PCM Signal
type PCM16bit struct {
	PCM
}

func NewPCM16bit(sampleRate uint32, Data []byte) PCM16bit {
	return PCM16bit{NewPCM(sampleRate, Data)}
}

func (s PCM16bit) property(p x) y {
	index := int(p/s.samplePeriod) * 2
	if index < 0 || index >= len(s.Data)-1 {
		return 0
	}
	return decodePCM16bit(s.Data[index], s.Data[index+1])
}

func encodePCM16bit(v y) (byte, byte) {
	return byte(v >> (yBits - 16)), byte(v >> (yBits - 8))
}

func decodePCM16bit(b1, b2 byte) y {
	return y(b1) << (yBits-16)|y(b2) << (yBits-8)
}

func (s PCM16bit) Encode(w io.Writer) {
	Encode(w, 2, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (s PCM16bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-2) / 2
}

func (s PCM16bit) Split(p x) (PCM16bit, PCM16bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 2)
	return PCM16bit{head}, PCM16bit{tail}
}

// 24 bit PCM Signal
type PCM24bit struct {
	PCM
}

func NewPCM24bit(sampleRate uint32, Data []byte) PCM24bit {
	return PCM24bit{NewPCM(sampleRate, Data)}
}

func (s PCM24bit) property(p x) y {
	index := int(p/s.samplePeriod) * 3
	if index < 0 || index >= len(s.Data)-2 {
		return 0
	}
	return decodePCM24bit(s.Data[index], s.Data[index+1], s.Data[index+2])
}
func encodePCM24bit(v y) (byte, byte, byte) {
	return byte(v >> (yBits - 24)), byte(v >> (yBits - 16)), byte(v >> (yBits - 8))
}
func decodePCM24bit(b1, b2, b3 byte) y {
	return y(b1) << (yBits-24)|y(b2) << (yBits-16)|y(b3) << (yBits-8)
}

func (s PCM24bit) Encode(w io.Writer) {
	Encode(w, 3, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (s PCM24bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-3) / 3
}

func (s PCM24bit) Split(p x) (PCM24bit, PCM24bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 3)
	return PCM24bit{head}, PCM24bit{tail}
}

// 32 bit PCM Signal
type PCM32bit struct {
	PCM
}

func NewPCM32bit(sampleRate uint32, Data []byte) PCM32bit {
	return PCM32bit{NewPCM(sampleRate, Data)}
}

func (s PCM32bit) property(p x) y {
	index := int(p/s.samplePeriod) * 4
	if index < 0 || index >= len(s.Data)-3 {
		return 0
	}
	return decodePCM32bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3])
}
func encodePCM32bit(v y) (byte, byte, byte, byte) {
	return byte(v >> (yBits - 32)), byte(v >> (yBits - 24)), byte(v >> (yBits - 16)), byte(v >> (yBits - 8))
}
func decodePCM32bit(b1, b2, b3, b4 byte) y {
	return y(b1) << (yBits-32)|y(b2) << (yBits-24)|y(b3) << (yBits-16)|y(b4) << (yBits-8)
}

func (s PCM32bit) Encode(w io.Writer) {
	Encode(w, 4, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (s PCM32bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-4) / 4
}

func (s PCM32bit) Split(p x) (PCM32bit, PCM32bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 4)
	return PCM32bit{head}, PCM32bit{tail}
}

// 48 bit PCM Signal
type PCM48bit struct {
	PCM
}

func NewPCM48bit(sampleRate uint32, Data []byte) PCM48bit {
	return PCM48bit{NewPCM(sampleRate, Data)}
}

func (s PCM48bit) property(p x) y {
	index := int(p/s.samplePeriod) * 6
	if index < 0 || index >= len(s.Data)-5 {
		return 0
	}
	return decodePCM48bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3], s.Data[index+4], s.Data[index+5])
}
func encodePCM48bit(v y) (byte, byte, byte, byte, byte, byte) {
	return byte(v >> (yBits - 48)), byte(v >> (yBits - 40)), byte(v >> (yBits - 32)), byte(v >> (yBits - 24)), byte(v >> (yBits - 16)), byte(v >> (yBits - 8))
}
func decodePCM48bit(b1, b2, b3, b4, b5, b6 byte) y {
	return y(b1) << (yBits-48)|y(b2) << (yBits-40)|y(b3) << (yBits-32)|y(b4) << (yBits-24)|y(b5) << (yBits-16)|y(b6) << (yBits-8)
}

func (s PCM48bit) Encode(w io.Writer) {
	Encode(w, 6, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (s PCM48bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-6) / 6
}

func (s PCM48bit) Split(p x) (PCM48bit, PCM48bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 6)
	return PCM48bit{head}, PCM48bit{tail}
}

// 64 bit PCM Signal
type PCM64bit struct {
	PCM
}

func NewPCM64bit(sampleRate uint32, Data []byte) PCM64bit {
	return PCM64bit{NewPCM(sampleRate, Data)}
}

func (s PCM64bit) property(p x) y {
	index := int(p/s.samplePeriod) * 8
	if index < 0 || index >= len(s.Data)-7 {
		return 0
	}
	return decodePCM64bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3], s.Data[index+4], s.Data[index+5], s.Data[index+6], s.Data[index+7])
}
func encodePCM64bit(v y) (byte, byte, byte, byte, byte, byte, byte, byte) {
	return byte(v >> (yBits - 64)), byte(v >> (yBits - 56)),byte(v >> (yBits - 48)), byte(v >> (yBits - 40)), byte(v >> (yBits - 32)), byte(v >> (yBits - 24)), byte(v >> (yBits - 16)), byte(v >> (yBits - 8))
}
func decodePCM64bit(b1, b2, b3, b4, b5, b6 , b7, b8 byte) y {
	return y(b1) << (yBits-64)|y(b1) << (yBits-56)|y(b1) << (yBits-48)|y(b2) << (yBits-40)|y(b3) << (yBits-32)|y(b4) << (yBits-24)|y(b5) << (yBits-16)|y(b6) << (yBits-8)
}

func (s PCM64bit) Encode(w io.Writer) {
	Encode(w, 8, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (s PCM64bit) MaxX() x {
	return s.PCM.samplePeriod * x(len(s.PCM.Data)-8) / 8
}

func (s PCM64bit) Split(p x) (PCM64bit, PCM64bit) {
	head, tail := s.PCM.Split(uint32(p/s.PCM.samplePeriod)+1, 8)
	return PCM64bit{head}, PCM64bit{tail}
}

// make a PeriodicLimitedSignal by sampling from another Signal, using provided parameters.
func NewPCMSignal(s Signal, length x, sampleRate uint32, sampleBytes uint8) PeriodicLimitedSignal {
	out, in := io.Pipe()
	go func() {
		Encode(in, sampleBytes, sampleRate, length, s)
		in.Close()
	}()
	channels, _ := Decode(out)
	out.Close()
	return channels[0]
}

