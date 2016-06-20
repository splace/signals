package signals

import (
	"fmt"
	"os"
	"testing"
	"net"
	"net/url"
)

func TestStreamSave(t *testing.T) {
	file, err := os.Create("./test output/stream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	//s:=Wav{URL:"http://www.nch.com.au/acm/8k8bitpcm.wav"}
	s,err:=NewWav("http://localhost:8086/wavs/s16/4.wav?f=8000")
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
		t.Error(err)
	}
	fmt.Printf("%#v\n",s)
	Encode(file, 2, 8000, unitX/3, s)
}


