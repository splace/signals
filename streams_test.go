package signals

import (
	"os"
	"testing"
	"net"
	"net/url"
)
const testURL="http://www.nch.com.au/acm/8k8bitpcm.wav"
// const testURL="http://localhost:8086/wavs/s16/4.wav?f=8000"

func TestStreamsRemoteSave(t *testing.T) {
	s,err:=NewWave(testURL)
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(ue.Error()+se.Err.Error())
					}
				}
			}
		}
		t.Fatal(err)
	}
	file, err := os.Create("./test output/remoteStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, s)
}

func TestStreamsLocalSave(t *testing.T) {
	s,err:=NewWave("http://localhost:8086/wavs/s16/4.wav?f=8000")
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(ue.Error()+se.Err.Error())
					}
				}
			}
		}
		t.Fatal(err)
	}
	file, err := os.Create("./test output/localStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, s)
}


func TestStreamsLocalRampUpSave(t *testing.T) {
	fs:=Modulated{&Wave{URL:"http://localhost:8086/wavs/s16/4.wav?f=8000"},RampUp{unitX}}
	file, err := os.Create("./test output/localFadeInStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, fs)
}

func TestStreamsGOBSaveLoadWave(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/wave.gob"); err != nil {panic(err)}else{defer file.Close()}
	m := Modulated{&Wave{URL:"http://localhost:8086/wavs/s16/4.wav?f=8000"},RampUp{unitX}}
	if err := Save(file,m); err != nil {
		panic(err)
	}
	file.Close()

	if file, err = os.Open("./test output/wave.gob"); err != nil {
		panic(err)
	}
	defer file.Close()

	s,err := Load(file)
	if err != nil {
		panic(err)
	}
	file, err = os.Create("./test output/gobFadeInStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, s)

}


