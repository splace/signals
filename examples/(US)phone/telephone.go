package main

import (
	. "../../../signals"
	"os"
)

func Save(file string,s Function){
	wavFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile,s,UnitX*4,44100,2)
}

const ms=UnitX/1000
/*
Audible Ring Tone is 440 Hz and 480 Hz for 2 seconds on and 4 seconds off
ReceiverOffHook is 1400 Hz, 2060 Hz, 2450 Hz, and 2600 Hz at 0 dBm0/frequency on and off every .1 second
No Such Number is 200 to 400 Hz modulated at 1 Hz, interrupted every 6 seconds for .5 seconds.
Line Busy Tone is 480 Hz and 630 Hz that is on and off every .5 seconds.
*/

func main(){
	Save("AudibleRingTone.wav",Looped{Multiplex{Pulse{UnitX*2},Stack{Sine{UnitX/440},Sine{UnitX/480}}},UnitX*6})
	Save("ReceiverOffHookTone.wav",Multiplex{Looped{Pulse{100*ms},200*ms}, Stack{Sine{UnitX / 1400},Sine{UnitX / 2060}, Sine{UnitX / 2450}, Sine{UnitX / 2600}}})
	Save("NoSuchNumberTone.wav",Stack{Sine{UnitX/200},Sine{UnitX/400}})
	Save("LineBusyTone.wav",Multiplex{Looped{Pulse{UnitX / 4},UnitX / 2}, Stack{Sine{UnitX/480},Sine{UnitX/630}}})

}

