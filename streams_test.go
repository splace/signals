package signals

import (
	"os"
	"testing"
	"net"
	"net/url"
)

func TestStreamSave(t *testing.T) {
	s,err:=NewWav("http://www.nch.com.au/acm/8k8bitpcm.wav")
	//s,err:=NewWav("http://localhost:8086/wavs/s16/4.wav?f=8000")
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
	file, err := os.Create("./test output/stream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	Encode(file, 1, 8000, unitX*3, *s)
}


