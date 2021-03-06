package signals

import (
	"testing"
	"os"
	"bytes"
	//"fmt"
)

func TestPCMscale(t *testing.T) {

	if decodePCM8bit(0x80)!=Y(0){t.Error(decodePCM8bit(0x80))}
	if decodePCM8bit(0xFF)+decodePCM8bit(0x81)-1!=Y(1){t.Error(Y(-1),decodePCM8bit(0xFF)+decodePCM8bit(0x81)-1)}
	if decodePCM8bit(0x00)+1!=Y(-1){t.Error(uint64(Y(-1)),uint64(decodePCM8bit(0x00))+1)}
	
	if encodePCM8bit(Y(0))!=0x80{t.Error(encodePCM8bit(Y(0)))}
	if encodePCM8bit(Y(1))!=0xFF{t.Error(encodePCM8bit(Y(1)))}
	if encodePCM8bit(Y(-1))!=0x00{t.Error(encodePCM8bit(Y(-1)))}

//	if Y(0)!= decodePCM16bit(encodePCM16bit(Y(0))) {t.Error(decodePCM16bit(encodePCM16bit(Y(0))))}
//	if Y(1)!= decodePCM16bit(encodePCM16bit(Y(1))) {t.Error(decodePCM16bit(encodePCM16bit(Y(1))))}
//	if Y(-1)!= decodePCM16bit(encodePCM16bit(Y(-1))) {t.Error(decodePCM16bit(encodePCM16bit(Y(-1))))}
//
//	if Y(0)!= decodePCM32bit(encodePCM32bit(Y(0))) {t.Error(decodePCM32bit(encodePCM32bit(Y(0))))}
//	if Y(1)!= decodePCM32bit(encodePCM32bit(Y(1))) {t.Error(decodePCM32bit(encodePCM32bit(Y(1))))}
//	if Y(-1)!= decodePCM32bit(encodePCM32bit(Y(-1))) {t.Error(decodePCM32bit(encodePCM32bit(Y(-1))))}

	
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


func TestPCMSaveLoad(t *testing.T) {
	pcm:=NewPCM(22050, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100})
	err:=pcm.SaveTo("./test output/pcm")
	if err!=nil {t.Error(err)}
	err=os.Remove("./test output/44100/pcm.pcm")
	if err!=nil {t.Error(err)}
	err=os.Remove("./test output/44100")
	if err!=nil {t.Error(err)}
	err=os.Rename("./test output/22050","./test output/44100")
	if err!=nil {t.Error(err)}
	var p PCM
	err=LoadPCM("./test output/44100/pcm",&p)
	if err!=nil {t.Error(err)}
	if p.samplePeriod != X(1 / float32(44100)) {t.Error("wrong sample rate")}
}


func TestPCMXSaveLoad(t *testing.T) {
	pcmx:=PCM16bit{NewPCM(22050, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100, 0, 110, 0, 120})}
	err:=pcmx.SaveTo("./test output/16bit/pcm")
	if err!=nil {t.Error(err)}
	var p PCM
	err=LoadPCM("./test output/16bit/22050/pcm",&p)
	if err!=nil {t.Error(err)}
	pcmx2:=PCM16bit{p}
	if !bytes.Equal(pcmx.PCM.Data,pcmx2.PCM.Data){t.Fail()}
	
//	if fmt.Sprintf("%#v", s) != fmt.Sprintf("%#v", m) {
//		t.Errorf("%#v != %#v", s, m)
//	}

}

func TestPCMXSaveLoadAny(t *testing.T) {
	pcmx:=PCM8bit{NewPCM(22050, []byte{0, 0, 0, 10, 0, 20, 0, 30, 0, 40, 0, 50, 0, 60, 0, 70, 0, 80, 0, 90, 0, 100, 0, 110, 0, 120})}
	err:=pcmx.SaveTo("./test output/8bit/pcm")
	if err!=nil {t.Error(err)}
	var p PCM
	err=LoadPCM("./test output/8bit/pcm",&p)
	if err!=nil {t.Error(err)}
	pcmx2:=PCM8bit{p}
	if !bytes.Equal(pcmx.PCM.Data,pcmx2.PCM.Data){t.Fail()}
	
//	if fmt.Sprintf("%#v", s) != fmt.Sprintf("%#v", m) {
//		t.Errorf("%#v != %#v", s, m)
//	}

}


var mb byte
var mx y

func BenchmarkPCM8bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		mb=encodePCM8bit(yv)
	}
}

func BenchmarkPCM8bitDecode(b *testing.B) {
	by:=byte(0)
	for i := 0; i < b.N; i++ {
		mx=decodePCM8bit(by)
	}
}


func BenchmarkPCM16bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		mb,mb=encodePCM16bit(yv)
	}
}
func BenchmarkPCM16bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	for i := 0; i < b.N; i++ {
		mx=decodePCM16bit(by1,by2)
	}
}


func BenchmarkPCM24bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		mb,mb,mb=encodePCM24bit(yv)
	}
}
func BenchmarkPCM24bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	by3:=byte(0)
	for i := 0; i < b.N; i++ {
		mx=decodePCM24bit(by1,by2,by3)
	}
}

func BenchmarkPCM32bitEncode(b *testing.B) {
	yv:=Y(0)
	for i := 0; i < b.N; i++ {
		mb,mb,mb,mb=encodePCM32bit(yv)
	}
}
func BenchmarkPCM32bitDecode(b *testing.B) {
	by1:=byte(0)
	by2:=byte(0)
	by3:=byte(0)
	by4:=byte(0)
	for i := 0; i < b.N; i++ {
		mx=decodePCM32bit(by1,by2,by3,by4)
	}
}



/*  Hal3 Sun Jun 19 00:53:10 BST 2016  go version go1.7beta2 linux/amd64

BenchmarkPCM8bitEncode-2    	2000000000	         1.16 ns/op
BenchmarkPCM8bitDecode-2    	2000000000	         0.77 ns/op
BenchmarkPCM16bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM16bitDecode-2   	2000000000	         1.38 ns/op
BenchmarkPCM24bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM24bitDecode-2   	2000000000	         1.73 ns/op
BenchmarkPCM32bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM32bitDecode-2   	1000000000	         2.31 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/signals	18.388s
Sun Jun 19 00:53:30 BST 2016 */
/*  Hal3 Sun Jun 19 00:53:41 BST 2016 go version go1.5.1 linux/amd64
PASS
BenchmarkPCM8bitEncode-2 	2000000000	         0.77 ns/op
BenchmarkPCM8bitDecode-2 	2000000000	         1.17 ns/op
BenchmarkPCM16bitEncode-2	2000000000	         1.16 ns/op
BenchmarkPCM16bitDecode-2	2000000000	         1.22 ns/op
BenchmarkPCM24bitEncode-2	2000000000	         1.93 ns/op
BenchmarkPCM24bitDecode-2	2000000000	         1.60 ns/op
BenchmarkPCM32bitEncode-2	2000000000	         1.94 ns/op
BenchmarkPCM32bitDecode-2	2000000000	         1.93 ns/op
ok  	_/home/simon/Dropbox/github/working/signals	24.686s
Sun Jun 19 00:54:08 BST 2016 */
/*  Hal3 Sun Jun 19 17:27:39 BST 2016 go version go1.5.1 linux/amd64
PASS
BenchmarkPCM8bitEncode-2 	2000000000	         0.77 ns/op
BenchmarkPCM8bitDecode-2 	2000000000	         1.18 ns/op
BenchmarkPCM16bitEncode-2	2000000000	         1.16 ns/op
BenchmarkPCM16bitDecode-2	2000000000	         1.54 ns/op
BenchmarkPCM24bitEncode-2	2000000000	         1.93 ns/op
BenchmarkPCM24bitDecode-2	1000000000	         1.98 ns/op
BenchmarkPCM32bitEncode-2	2000000000	         1.94 ns/op
BenchmarkPCM32bitDecode-2	1000000000	         2.52 ns/op
ok  	_/home/simon/Dropbox/github/working/signals	22.938s
Sun Jun 19 17:28:04 BST 2016 */
/*  Hal3 Sun Jun 19 17:41:12 BST 2016  go version go1.7beta2 linux/amd64

BenchmarkPCM8bitEncode-2    	2000000000	         1.16 ns/op
BenchmarkPCM8bitDecode-2    	2000000000	         0.77 ns/op
BenchmarkPCM16bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM16bitDecode-2   	2000000000	         0.77 ns/op
BenchmarkPCM24bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM24bitDecode-2   	2000000000	         0.77 ns/op
BenchmarkPCM32bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM32bitDecode-2   	2000000000	         0.77 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/signals	14.189s
Sun Jun 19 17:41:27 BST 2016 */



/*  Hal3 Tue 16 Aug 00:42:15 BST 2016 go version go1.6.2 linux/amd64
PASS
BenchmarkPCM8bitEncode-2 	2000000000	         0.77 ns/op
BenchmarkPCM8bitDecode-2 	2000000000	         1.16 ns/op
BenchmarkPCM16bitEncode-2	2000000000	         1.16 ns/op
BenchmarkPCM16bitDecode-2	2000000000	         1.16 ns/op
BenchmarkPCM24bitEncode-2	2000000000	         1.93 ns/op
BenchmarkPCM24bitDecode-2	2000000000	         1.54 ns/op
BenchmarkPCM32bitEncode-2	2000000000	         1.98 ns/op
BenchmarkPCM32bitDecode-2	2000000000	         1.93 ns/op
ok  	_/home/simon/Dropbox/github/working/signals	24.535s
Tue 16 Aug 00:42:42 BST 2016 */
/*  Hal3 Tue 16 Aug 00:42:55 BST 2016  go version go1.7 linux/amd64

BenchmarkPCM8bitEncode-2    	2000000000	         1.16 ns/op
BenchmarkPCM8bitDecode-2    	2000000000	         0.77 ns/op
BenchmarkPCM16bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM16bitDecode-2   	2000000000	         1.20 ns/op
BenchmarkPCM24bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM24bitDecode-2   	2000000000	         1.68 ns/op
BenchmarkPCM32bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM32bitDecode-2   	1000000000	         2.06 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/signals	17.702s
Tue 16 Aug 00:43:14 BST 2016 */
/*  Hal3 Tue 31 Oct 19:46:06 GMT 2017  go version go1.9.1 linux/amd64

goos: linux
goarch: amd64
BenchmarkPCM8bitEncode-2    	2000000000	         1.16 ns/op
BenchmarkPCM8bitDecode-2    	2000000000	         0.77 ns/op
BenchmarkPCM16bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM16bitDecode-2   	2000000000	         0.77 ns/op
BenchmarkPCM24bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM24bitDecode-2   	2000000000	         0.77 ns/op
BenchmarkPCM32bitEncode-2   	2000000000	         0.83 ns/op
BenchmarkPCM32bitDecode-2   	2000000000	         0.77 ns/op
PASS
ok  	_/home/simon/Dropbox/github/working/signals	14.206s
Tue 31 Oct 19:46:29 GMT 2017
*/

