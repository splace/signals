package main

import (
	. "github.com/splace/signals" //"../../../signals" // 
	"os"
)

func Save(file string,s PeriodicFunction){
	wavFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// whole number of cycles or at least a seconds worth
	if s.Period()>X(1){
		Encode(wavFile,s,s.Period(),44100,2)
	}else{
		Encode(wavFile,s,s.Period()*(X(1)/s.Period()),44100,2)
	}
}

/*
Audible Ring Tone is 440 Hz and 480 Hz for 2 seconds on and 4 seconds off
ReceiverOffHook is 1400 Hz, 2060 Hz, 2450 Hz, and 2600 Hz at 0 dBm0/frequency on and off every .1 second
No Such Number is 200 to 400 Hz modulated at 1 Hz, interrupted every 6 seconds for .5 seconds.
Line Busy Tone is 480 Hz and 630 Hz that is on and off every .5 seconds.
*/

func main(){
	Save("AudibleRingTone.wav",Looped{Multiplex{Pulse{X(2)},Stack{Sine{X(1.0/440)},Sine{X(1.0/480)}}},X(6)})
	Save("ReceiverOffHookTone.wav",Multiplex{Looped{Pulse{X(.1)},X(.2)}, Stack{Sine{X(1.0/1400)},Sine{X(1.0/2060)}, Sine{X(1.0/2450)}, Sine{X(1.0/2600)}}})
	Save("NoSuchNumberTone.wav",Stack{Sine{X(1.0/200)},Sine{X(1.0/400)}})
	Save("LineBusyTone.wav",Multiplex{Looped{Pulse{X(.25)},X(0.5)}, Stack{Sine{X(1.0/480)},Sine{X(1/630)}}})

}



