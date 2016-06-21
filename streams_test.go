package signals

import (
	"os"
	"testing"
	"net"
	"net/url"
)

func TestStreamsRemoteSave(t *testing.T) {
	s,err:=NewWave("http://www.nch.com.au/acm/8k8bitpcm.wav")
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(se.Err.Error())
					}
				}
			}
		}
		t.Fatal(err)
	}
	t.Logf("%v\n",s)
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
						t.Skip(se.Err.Error())
					}
				}
			}
		}
		t.Fatal(err)
	}
	t.Logf("%v\n",s)
	file, err := os.Create("./test output/localStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, s)
}


func TestStreamsLocalRampUpSave(t *testing.T) {
	s,err:=NewWave("http://localhost:8086/wavs/s16/4.wav?f=8000")
	if err!=nil{
		if ue,ok:=err.(*url.Error);ok {
			if oe,ok:=ue.Err.(*net.OpError);ok{
				if se,ok:=oe.Err.(*os.SyscallError);ok{
					if se.Err.Error()=="connection refused"{
						t.Skip(se.Err.Error())
					}
				}
			}
		}
		t.Fatal(err)
	}
	t.Logf("%v\n",s)
	fs:=Modulated{s,RampUp{unitX}}
	file, err := os.Create("./test output/localFadeInStream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, fs)
}


