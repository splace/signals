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

/*  Hal3 Thu May 26 00:23:26 BST 2016 go version go1.5.1 linux/amd64
Thu May 26 00:23:26 BST 2016 */
/*  Hal3 Thu May 26 00:23:44 BST 2016 go version go1.5.1 linux/amd64
Thu May 26 00:23:48 BST 2016 */
/*  Hal3 Thu May 26 00:24:33 BST 2016 go version go1.5.1 linux/amd64
Thu May 26 00:24:37 BST 2016 */

