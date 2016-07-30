package signals

import (
	"os"
	"testing"
	"net"
	"net/url"
)

const testLocalURL="http://localhost:8086/wavs/s16/4.wav?f=8000"

const testRemoteURL="http://www.nch.com.au/acm/8k8bitpcm.wav"  //const testRemoteURL="http://localhost:8086/wavs/s16/4.wav?f=8000"  //

const testDataURL="data:audio/x-wav;base64,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"

const testFileURL="file:///home/simon/Dropbox/github/working/signals/middlec.wav"
const testGOBFile="/home/simon/Dropbox/github/working/signals/test output/RingingTone"
func TestStreamsRemoteSave(t *testing.T) {
	s,err:=NewWave(testRemoteURL)
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(ue.Error())
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
	s,err:=NewWave(testLocalURL)
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(ue.Error())
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
	fs:=Modulated{&Wave{URL:testLocalURL},RampUp{unitX}}
	file, err := os.Create("./test output/localFadeInStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, fs)
}


func TestStreamsSaveDataURL(t *testing.T) {
	fs:=&Wave{URL:testDataURL}
	file, err := os.Create("./test output/dataURL.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, fs)
}


func TestStreamsSaveFileURL(t *testing.T) {
	fs:=&Wave{URL:testFileURL}
	file, err := os.Create("./test output/fileURL.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, fs)
}


func TestStreamsSaveGOBFileURL(t *testing.T) {
	err := SaveGOB(testGOBFile,Looped{Modulated{Pulse{unitX}, Looped{Pulse{unitX * 4 / 10}, unitX * 6 / 10}, Stacked{Sine{unitX / 450}, Sine{unitX / 400}}}, unitX * 3})
	if err != nil { t.Error(err)}
	fs:=&Wave{URL:"file://"+testGOBFile+".gob"}
	file, err := os.Create("./test output/GOBfileURL.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 2, 22050, unitX*6, fs)
}

const testPCMFile="/home/simon/Dropbox/github/working/signals/test output/16bit/22050/pcm.pcm"

func TestStreamsSavePCMFileURL(t *testing.T) {
	fs:=&Wave{URL:"file://"+testPCMFile}
	file, err := os.Create("./test output/PCMfileURL.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 2, 22050, unitX/2205, fs)
}



