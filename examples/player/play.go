package main

import (
	"os/exec"
	"io"
	"bytes"
)

import . "github.com/splace/signals"

var OneSecond = X(1)

func play(s Function) {
	cmd := exec.Command("aplay","-f","S16","-r","44100")
	out, in := io.Pipe()
	go func() {
		io.Copy(in, bytes.NewReader(NewPCMFunction(s, OneSecond*3, 44100, 2).(PCM16bit).PCM.data))
		in.Close()
	}()
	cmd.Stdin=out 
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main(){
	play(Looped{Modulated{Pulse{OneSecond}, Looped{Pulse{OneSecond*4/10}, OneSecond*6/10}, Stack{Sine{OneSecond/450},Sine{OneSecond/400}}}, OneSecond*3})
}


/*  Hal3 Fri Apr 29 00:28:36 BST 2016 go version go1.5.1 linux/amd64
Fri Apr 29 00:28:36 BST 2016 */
/*  Hal3 Fri Apr 29 00:31:06 BST 2016 go version go1.5.1 linux/amd64
Fri Apr 29 00:31:06 BST 2016 */

