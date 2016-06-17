package signals

import (
	"testing"
	"os"
	//"fmt"
)

func TestPCMscale(t *testing.T) {

	if PCM8bitDecode(0x80)!=Y(0){t.Error(PCM8bitDecode(0x80))}
	if PCM8bitDecode(0xFF)+PCM8bitDecode(0x81)-1!=Y(1){t.Error(Y(-1),PCM8bitDecode(0xFF)+PCM8bitDecode(0x81)-1)}
	if PCM8bitDecode(0x00)+1!=Y(-1){t.Error(uint64(Y(-1)),uint64(PCM8bitDecode(0x00))+1)}
	
	if PCM8bitEncode(Y(0))!=0x80{t.Error(PCM8bitEncode(Y(0)))}
	if PCM8bitEncode(Y(1))!=0xFF{t.Error(PCM8bitEncode(Y(1)))}
	if PCM8bitEncode(Y(-1))!=0x00{t.Error(PCM8bitEncode(Y(-1)))}
	
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



