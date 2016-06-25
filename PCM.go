package signals

import (
	"io"
)

// PCM is the state and behaviour common to all PCM, it doesn't include encoding information, so cannot return a property and so its not a Signal.
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

// 8 bit PCM Signal.
// unlike the other precisions of PCM, that use signed data, 8bit uses un-signed. (the default OpenAL and wave file representation for 8bit precision.)
type PCM8bit struct {
	PCM
}

func NewPCM8bit(sampleRate uint32, Data []byte) PCM8bit {
	return PCM8bit{NewPCM(sampleRate, Data)}
}

func (s PCM8bit) property(offset x) y {
	index := int(offset / s.samplePeriod)
	if index < 0 || index >= len(s.Data){
		return 0
	}
	return decodePCM8bit(s.Data[index])
}


func encodePCM8bit(y y) byte {
	return byte(y>>(yBits-8)) + 128
}

func decodePCM8bit(b byte) y {
	return y(b-128) << (yBits-8)
}

func (p PCM8bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-1)
}

func (s PCM8bit) Encode(w io.Writer) {
	Encode(w, 1, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM8bit) Split(position x) (PCM8bit, PCM8bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 1)
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
	if index < 0 || index >= len(s.Data)-1 {
		return 0
	}
	return decodePCM16bit(s.Data[index], s.Data[index+1])
}

func encodePCM16bit(y y) (byte, byte) {
	return byte(y >> (yBits - 16)), byte(y >> (yBits - 8))
}

func decodePCM16bit(b1, b2 byte) y {
	return y(b1) << (yBits-16)|y(b2) << (yBits-8)
}

func (s PCM16bit) Encode(w io.Writer) {
	Encode(w, 2, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM16bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-2) / 2
}

func (p PCM16bit) Split(position x) (PCM16bit, PCM16bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 2)
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
	if index < 0 || index >= len(s.Data)-2 {
		return 0
	}
	return decodePCM24bit(s.Data[index], s.Data[index+1], s.Data[index+2])
}
func encodePCM24bit(y y) (byte, byte, byte) {
	return byte(y >> (yBits - 24)), byte(y >> (yBits - 16)), byte(y >> (yBits - 8))
}
func decodePCM24bit(b1, b2, b3 byte) y {
	return y(b1) << (yBits-24)|y(b2) << (yBits-16)|y(b3) << (yBits-8)
}

func (s PCM24bit) Encode(w io.Writer) {
	Encode(w, 3, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM24bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-3) / 3
}

func (p PCM24bit) Split(position x) (PCM24bit, PCM24bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 3)
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
	if index < 0 || index >= len(s.Data)-3 {
		return 0
	}
	return decodePCM32bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3])
}
func encodePCM32bit(y y) (byte, byte, byte, byte) {
	return byte(y >> (yBits - 32)), byte(y >> (yBits - 24)), byte(y >> (yBits - 16)), byte(y >> (yBits - 8))
}
func decodePCM32bit(b1, b2, b3, b4 byte) y {
	return y(b1) << (yBits-32)|y(b2) << (yBits-24)|y(b3) << (yBits-16)|y(b4) << (yBits-8)
}

func (s PCM32bit) Encode(w io.Writer) {
	Encode(w, 4, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM32bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-4) / 4
}

func (p PCM32bit) Split(position x) (PCM32bit, PCM32bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 4)
	return PCM32bit{head}, PCM32bit{tail}
}

// 48 bit PCM Signal
type PCM48bit struct {
	PCM
}

func NewPCM48bit(sampleRate uint32, Data []byte) PCM48bit {
	return PCM48bit{NewPCM(sampleRate, Data)}
}

func (s PCM48bit) property(offset x) y {
	index := int(offset/s.samplePeriod) * 6
	if index < 0 || index >= len(s.Data)-5 {
		return 0
	}
	return decodePCM48bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3], s.Data[index+4], s.Data[index+5])
}
func encodePCM48bit(y y) (byte, byte, byte, byte, byte, byte) {
	return byte(y >> (yBits - 48)), byte(y >> (yBits - 40)), byte(y >> (yBits - 32)), byte(y >> (yBits - 24)), byte(y >> (yBits - 16)), byte(y >> (yBits - 8))
}
func decodePCM48bit(b1, b2, b3, b4, b5, b6 byte) y {
	return y(b1) << (yBits-48)|y(b2) << (yBits-40)|y(b3) << (yBits-32)|y(b4) << (yBits-24)|y(b5) << (yBits-16)|y(b6) << (yBits-8)
}

func (s PCM48bit) Encode(w io.Writer) {
	Encode(w, 6, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM48bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-6) / 6
}

func (p PCM48bit) Split(position x) (PCM48bit, PCM48bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 6)
	return PCM48bit{head}, PCM48bit{tail}
}

// 64 bit PCM Signal
type PCM64bit struct {
	PCM
}

func NewPCM64bit(sampleRate uint32, Data []byte) PCM64bit {
	return PCM64bit{NewPCM(sampleRate, Data)}
}

func (s PCM64bit) property(offset x) y {
	index := int(offset/s.samplePeriod) * 8
	if index < 0 || index >= len(s.Data)-7 {
		return 0
	}
	return decodePCM64bit(s.Data[index], s.Data[index+1], s.Data[index+2], s.Data[index+3], s.Data[index+4], s.Data[index+5], s.Data[index+6], s.Data[index+7])
}
func encodePCM64bit(y y) (byte, byte, byte, byte, byte, byte, byte, byte) {
	return byte(y >> (yBits - 64)), byte(y >> (yBits - 56)),byte(y >> (yBits - 48)), byte(y >> (yBits - 40)), byte(y >> (yBits - 32)), byte(y >> (yBits - 24)), byte(y >> (yBits - 16)), byte(y >> (yBits - 8))
}
func decodePCM64bit(b1, b2, b3, b4, b5, b6 , b7, b8 byte) y {
	return y(b1) << (yBits-64)|y(b1) << (yBits-56)|y(b1) << (yBits-48)|y(b2) << (yBits-40)|y(b3) << (yBits-32)|y(b4) << (yBits-24)|y(b5) << (yBits-16)|y(b6) << (yBits-8)
}

func (s PCM64bit) Encode(w io.Writer) {
	Encode(w, 8, uint32(unitX/s.Period()), s.MaxX(), s)
}
func (p PCM64bit) MaxX() x {
	return p.PCM.samplePeriod * x(len(p.PCM.Data)-8) / 8
}

func (p PCM64bit) Split(position x) (PCM64bit, PCM64bit) {
	head, tail := p.PCM.Split(uint32(position/p.PCM.samplePeriod)+1, 8)
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

