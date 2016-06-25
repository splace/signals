package signals

import (
	"os"
	"testing"
	"net"
	"net/url"
)

func TestCacheStreamsSave(t *testing.T) {
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
	//t.Logf("%v\n",s)
	fs:=Cached{s,make(map[x]y)}
	file, err := os.Create("./test output/cachedStream.wav")
	if err != nil {panic(err)}
	Encode(file, 1, 8000, unitX*3, fs)
	file.Close()

}



