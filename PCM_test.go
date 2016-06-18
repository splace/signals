package signals

import (
	"testing"
	"os"
	//"fmt"
)

func TestPCMscale(t *testing.T) {

	if decodePCM8bit(0x80)!=Y(0){t.Error(decodePCM8bit(0x80))}
	if decodePCM8bit(0xFF)+decodePCM8bit(0x81)-1!=Y(1){t.Error(Y(-1),decodePCM8bit(0xFF)+decodePCM8bit(0x81)-1)}
	if decodePCM8bit(0x00)+1!=Y(-1){t.Error(uint64(Y(-1)),uint64(decodePCM8bit(0x00))+1)}
	
	if encodePCM8bit(Y(0))!=0x80{t.Error(encodePCM8bit(Y(0)))}
	if encodePCM8bit(Y(1))!=0xFF{t.Error(encodePCM8bit(Y(1)))}
	if encodePCM8bit(Y(-1))!=0x00{t.Error(encodePCM8bit(Y(-1)))}
	
}

func TestPCMRaw(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/TestRaw.wav"); err != nil {panic(err)}else{defer file.Close()}
	PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})}.Encode(file)
}

func TestPCMSplit(t *testing.T) {
	var wavFileHead *os.File
	var wavFileTail *os.File
	var err error
	if wavFileHead, err = os.Create("./test output/TestSplitHead.wav"); err != nil {panic(err)}else{defer wavFileHead.Close()}
	if wavFileTail, err = os.Create("./test output/TestSplitTail.wav"); err != nil {panic(err)}else{defer wavFileTail.Close()}
	sh, st := PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})}.Split(X(1.01))
	sh.Encode(wavFileHead)
	st.Encode(wavFileTail)
}


func TestPCMEnocdeToShortLength(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/EnocdePCMToShortLength.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 2, 5, unitX,PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})})
}


func TestPCMEnocdeShiftedPCM(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/EnocdeShiftedPCM.wav"); err != nil {panic(err)}else{defer file.Close()}
	Encode(file, 2, 5, unitX/2,Shifted{PCM16bit{NewPCM(5, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})},unitX})
}

func BenchmarkPCM8bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		encodePCM8bit(yv)
	}
}
func BenchmarkPCM8bitDecode(b *testing.B) {
	by:=byte(0)
	for i := 0; i < b.N; i++ {
		decodePCM8bit(by)
	}
}


func BenchmarkPCM16bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		encodePCM16bit(yv)
	}
}
func BenchmarkPCM16bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	for i := 0; i < b.N; i++ {
		decodePCM16bit(by1,by2)
	}
}


func BenchmarkPCM24bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		encodePCM24bit(yv)
	}
}
func BenchmarkPCM24bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	by3:=byte(0)
	for i := 0; i < b.N; i++ {
		decodePCM24bit(by1,by2,by3)
	}
}

func BenchmarkPCM32bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		encodePCM32bit(yv)
	}
}
func BenchmarkPCM32bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	by3:=byte(0)
	by4:=byte(0)
	for i := 0; i < b.N; i++ {
		decodePCM32bit(by1,by2,by3,by4)
	}
}



/*  Hal3 Sat Jun 18 21:23:16 BST 2016 go version go1.5.1 linux/amd64
PASS
BenchmarkPCM8bitEncode-2 	2000000000	         0.81 ns/op
BenchmarkPCM8bitDecode-2 	2000000000	         0.91 ns/op
BenchmarkPCM16bitEncode-2	2000000000	         0.92 ns/op
BenchmarkPCM16bitDecode-2	2000000000	         1.20 ns/op
BenchmarkPCM24bitEncode-2	2000000000	         1.16 ns/op
BenchmarkPCM24bitDecode-2	2000000000	         1.55 ns/op
BenchmarkPCM32bitEncode-2	2000000000	         1.42 ns/op
BenchmarkPCM32bitDecode-2	2000000000	         1.95 ns/op
ok  	_/home/simon/Dropbox/github/working/signals	20.907s
Sat Jun 18 21:23:39 BST 2016 */

