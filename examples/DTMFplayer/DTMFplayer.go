// pipe command for converting characters into DTMF tone PCM data.
// example usage:  play text as tones.
// ./DTMFplayer[SYSV64].elf <<< "0123456789ABCD#*" | aplay -fs16 -r 16000
// note:
// "aplay" parameters should match sample rate and format compiled for.
package main

import (
	"io"
	"os"
	"bufio"
)



import . "github.com/splace/signals"

var length=X(.07)
var gap=X(.08)
var sampleRate uint32=16000

var Tones = map[rune]PCM{
	'0':NewPCMFunction( Stack{Sine{X(1.0/941)},Sine{X(1.0/1336)}},length, sampleRate,2).(PCM16bit).PCM,
	'1':NewPCMFunction( Stack{Sine{X(1.0/697)},Sine{X(1.0/1209)}},length, sampleRate,2).(PCM16bit).PCM,
	'2':NewPCMFunction( Stack{Sine{X(1.0/697)},Sine{X(1.0/1336)}},length, sampleRate,2).(PCM16bit).PCM,
	'3':NewPCMFunction( Stack{Sine{X(1.0/697)},Sine{X(1.0/1477)}},length, sampleRate,2).(PCM16bit).PCM,
	'4':NewPCMFunction( Stack{Sine{X(1.0/770)},Sine{X(1.0/1209)}},length, sampleRate,2).(PCM16bit).PCM,
	'5':NewPCMFunction( Stack{Sine{X(1.0/770)},Sine{X(1.0/1336)}},length, sampleRate,2).(PCM16bit).PCM,
	'6':NewPCMFunction( Stack{Sine{X(1.0/770)},Sine{X(1.0/1477)}},length, sampleRate,2).(PCM16bit).PCM,
	'7':NewPCMFunction( Stack{Sine{X(1.0/852)},Sine{X(1.0/1209)}},length, sampleRate,2).(PCM16bit).PCM,
	'8':NewPCMFunction( Stack{Sine{X(1.0/852)},Sine{X(1.0/1336)}},length, sampleRate,2).(PCM16bit).PCM,
	'9':NewPCMFunction( Stack{Sine{X(1.0/852)},Sine{X(1.0/1477)}},length, sampleRate,2).(PCM16bit).PCM,
	'A':NewPCMFunction( Stack{Sine{X(1.0/697)},Sine{X(1.0/1633)}},length, sampleRate,2).(PCM16bit).PCM,
	'B':NewPCMFunction( Stack{Sine{X(1.0/770)},Sine{X(1.0/1633)}},length, sampleRate,2).(PCM16bit).PCM,
	'C':NewPCMFunction( Stack{Sine{X(1.0/852)},Sine{X(1.0/1633)}},length, sampleRate,2).(PCM16bit).PCM,
	'D':NewPCMFunction( Stack{Sine{X(1.0/941)},Sine{X(1.0/1633)}},length, sampleRate,2).(PCM16bit).PCM,
	'*':NewPCMFunction( Stack{Sine{X(1.0/941)},Sine{X(1.0/1209)}},length, sampleRate,2).(PCM16bit).PCM,
	'#':NewPCMFunction( Stack{Sine{X(1.0/941)},Sine{X(1.0/1477)}},length, sampleRate,2).(PCM16bit).PCM,
}

var gapPCM=NewPCMFunction(Constant{0},gap, sampleRate,2).(PCM16bit).PCM

func main(){
	rr:=bufio.NewReader(os.Stdin)
	for ;;{
		Rune,_,err:=rr.ReadRune()
		if err==io.EOF {
			os.Stdin.Close()
			break
		}else if err!=nil{
			panic(err)
		}
		Tones[Rune].Encode(os.Stdout)
		gapPCM.Encode(os.Stdout)
	}
	os.Stdout.Close()
}




