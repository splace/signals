// generate a few standard telephone notification tones.
// length is a multiple of the repeat cycle, so play output repeatedly to get any length. 
package main

import (
	. "github.com/splace/signals" 
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
``On'' and ``off'' Dxs are in ms. The frequency is 400 Hz, except where noted.
   	  	On 	Off 	On 	Off 	Notes 	Audio sample
BT 	Busy tone 	375 	375 	  	  	  	[AU]
EET 	Equipment engaged tone 	400 	350 	225 	525 	1 	[AU]
RT 	Ringing tone 	400 	200 	400 	2000 	2 	[AU]
NU 	Number unobtainable 	Continuous 	  	[AU]
DT 	Dial tone 	Continuous 	4 	[AU]
Notes

 1   The amplitude of the 225ms tone is 6dB higher than that of the 400mS tone. This is specified (I'm reliably told) in BS 6305 (1992). I'm grateful to Nigel Roles <ngr@symbionics.co.uk> for pointing this out.
 2   Frequency: 400+450 Hz.
 4   Frequency: 350+450 Hz.

*/

func main(){
	Save("BusyTone.wav",Modulated{Looped{Pulse{OneSecond*375/1000},OneSecond*75/100}, Sine{OneSecond/400}})
	Save("EngagedTone.wav",Looped{Modulated{Composite{Modulated{Pulse{OneSecond*4/10},NewConstant(-6)},Shifted{Pulse{OneSecond*225/1000},OneSecond*75/100}}, Sine{OneSecond/400}}, OneSecond*15/10})
	Save("RingingTone.wav",Looped{Modulated{Pulse{OneSecond}, Looped{Pulse{OneSecond*4/10}, OneSecond*6/10}, Stack{Sine{OneSecond/450},Sine{OneSecond/400}}}, OneSecond*3})
	Save("NumberUnobtainableTone.wav",Sine{OneSecond/400})
	Save("dialTone.wav",Stack{Sine{OneSecond/450},Sine{OneSecond/350}})

}

