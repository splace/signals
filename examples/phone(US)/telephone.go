// generate a few standard US telephone notification tones.
// length is a multiple of cycle, so play wave files repeated to get any length. 
package main

import (
	. "github.com/splace/signals" //"../../../signals" // 
	"os"
)
var OneSecond = X(1)

func Save(file string,s PeriodicFunction){
	wavFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// one cycle or at least a seconds worth
	if s.Period()>OneSecond{
		Encode(wavFile,s,s.Period(),44100,2)
	}else{
		Encode(wavFile,s,s.Period()*(OneSecond/s.Period()),44100,2)
	}
}

/*
Audible Ring Tone is 440 Hz and 480 Hz for 2 seconds on and 4 seconds off
ReceiverOffHook is 1400 Hz, 2060 Hz, 2450 Hz, and 2600 Hz at 0 dBm0/frequency on and off every .1 second
No Such Number is 200 to 400 Hz modulated at 1 Hz, interrupted every 6 seconds for .5 seconds.
Line Busy Tone is 480 Hz and 630 Hz that is on and off every .5 seconds.
*/

func main(){
	Save("AudibleRingTone.wav",Looped{Modulated{Pulse{OneSecond*2},Stack{Sine{OneSecond/440},Sine{OneSecond/480}}},OneSecond*6})
	Save("ReceiverOffHookTone.wav",Modulated{Looped{Pulse{OneSecond/10},OneSecond/5}, Stack{Sine{OneSecond/1400},Sine{OneSecond/2060}, Sine{OneSecond/2450}, Sine{OneSecond/2600}}})
	Save("NoSuchNumberTone.wav",Stack{Sine{OneSecond/200},Sine{OneSecond/400}})
	Save("LineBusyTone.wav",Modulated{Looped{Pulse{OneSecond/4},OneSecond/2}, Stack{Sine{OneSecond/480},Sine{OneSecond/630}}})

}



