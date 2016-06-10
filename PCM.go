package signals

import (
	"io"
)

// PCM is the state and behaviour common to all PCM. Its not a Signal, specific PCM<<precison>> types embed this, and then are Signal's.
// the specific precision types, the Signals, return continuous property values that step from one PCM value to the next, Segmented could be used to get interpolated property values.
type PCM struct {
	samplePeriod x
	Data         []byte
}

// make a PCM type, from raw bytes.
func NewPCM(sampleRate uint32, Data []byte) PCM {
	return PCM{X(1 / float32(sampleRate)), Data}
}

func (p PCM) Period() x {
	return p.samplePeriod
}

// from a PCM return two new PCM's (with the same underlying data) from either side of a sample.
func (p PCM) Split(sample uint32, sampleBytes uint8) (head PCM, tail PCM) {
	copy := func(p PCM) PCM { return p }
	bytePosition := sample * uint32(sampleBytes)
	if bytePosition > uint32(len(p.Data)) {
		bytePosition = uint32(len(p.Data))
	}
	head, tail = p, copy(p)
	tail.Data = tail.Data[bytePosition:]
	head.Data = head.Data[:bytePosition]
	return
}

// 8 bit PCMSignal.
// unlike the other precisions of PCM, that use signed data, 8bit uses un-signed. (the default OpenAL and wave file representation for 8bit precision.)
type PCM8bit struct {
	PCM
}

func NewPCM8bit(sampleRate uint32, Data []byte) PCM8bit {
	return PCM8bit{NewPCM(sampleRate, Data)}
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

func (p PCM8bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-1)
}

func (p PCM8bit) Split(position x) (PCM8bit, PCM8bit) {
	head, tail := p.PCM.Split(uint32(uint64(len(p.PCM.Data))*uint64(position)/uint64(p.MaxX()))+1, 1)
	return PCM8bit{head}, PCM8bit{tail}
}

// 16 bit PCM Signal
type PCM16bit struct {
	PCM
}

func NewPCM16bit(sampleRate uint32, Data []byte) PCM16bit {
	return PCM16bit{NewPCM(sampleRate, Data)}
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
func (p PCM16bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-2) / 2
}

func (p PCM16bit) Split(position x) (PCM16bit, PCM16bit) {
	head, tail := p.PCM.Split(uint32(uint64(len(p.PCM.Data)/2)*uint64(position)/uint64(p.MaxX()))+1, 2)
	return PCM16bit{head}, PCM16bit{tail}
}

// 24 bit PCM Signal
type PCM24bit struct {
	PCM
}

func NewPCM24bit(sampleRate uint32, Data []byte) PCM24bit {
	return PCM24bit{NewPCM(sampleRate, Data)}
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
func (p PCM24bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-3) / 3
}

func (p PCM24bit) Split(position x) (PCM24bit, PCM24bit) {
	head, tail := p.PCM.Split(uint32(uint64(len(p.PCM.Data)/3)*uint64(position)/uint64(p.MaxX()))+1, 3)
	return PCM24bit{head}, PCM24bit{tail}
}

// 32 bit PCM Signal
type PCM32bit struct {
	PCM
}

func NewPCM32bit(sampleRate uint32, Data []byte) PCM32bit {
	return PCM32bit{NewPCM(sampleRate, Data)}
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
func (p PCM32bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-4) / 4
}

func (p PCM32bit) Split(position x) (PCM32bit, PCM32bit) {
	head, tail := p.PCM.Split(uint32(uint64(len(p.PCM.Data)/4)*uint64(position)/uint64(p.MaxX()))+1, 4)
	return PCM32bit{head}, PCM32bit{tail}
}

// make a PeriodicLimitedSignal by sampling from another Signal, using provided parameters.
func NewPCMSignal(s Signal, length x, sampleRate uint32, sampleBytes uint8) PeriodicLimitedSignal {
	out, in := io.Pipe()
	go func() {
		Encode(in, s, length, sampleRate, sampleBytes)
		in.Close()
	}()
	channels, _ := Decode(out)
	out.Close()
	return channels[0]
}

