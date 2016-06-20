package signals

import (
	"fmt"
	"os"
	"testing"
)

func TestStreamSave(t *testing.T) {
	file, err := os.Create("./test output/stream.wav")
	if err != nil {panic(err)}
	defer file.Close()
	//s:=Wav{URL:"http://www.nch.com.au/acm/8k8bitpcm.wav"}
	s,err:=NewWav("http://localhost:8086/wavs/s16/4.wav?f=8000")
	if err!=nil{panic(err)}
	fmt.Printf("%#v\n",s)
	Encode(file, 2, 8000, unitX/3, s)
}


