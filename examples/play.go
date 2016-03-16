package main

import (
	"os/exec"
	"io"
)

import . "github.com/splace/signals"

var OneSecond = X(1)

func play(s Function) {
	cmd := exec.Command("aplay","-fS16","-r 44100")
	out, in := io.Pipe()
	go func() {
		NewPCMFunction(s, OneSecond*3, 44100, 2).(PCM16bit).PCM.Encode(in)
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


