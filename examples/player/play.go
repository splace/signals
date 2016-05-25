// play (needed aplay) telephone ringing tone, one cycle.
package main

import (
	"os/exec"
	"io"
	"bytes"
)

import . "github.com/splace/signals"

var OneSecond = X(1)

func play(s Signal) {
	cmd := exec.Command("aplay","-f","S16","-r","44100")
	out, in := io.Pipe()
	go func() {
		io.Copy(in, bytes.NewReader(NewPCMSignal(s, OneSecond*3, 44100, 2).(PCM16bit).PCM.Data))
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

