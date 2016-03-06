// generate a few standard telephone notification signals
// multiples of cycle, so play wave files repeated to get any length. 
package main

import (
	. "github.com/splace/signals"  //"../../../signals" // 
	"os"
)

func Save(file string,s PeriodicFunction){
	wavFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	// one cycle or at least a seconds worth
	if s.Period()>X(1){
		Encode(wavFile,s,s.Period(),44100,2)
	}else{
		Encode(wavFile,s,s.Period()*(X(1)/s.Period()),44100,2)
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
	Save("BusyTone.wav",Multiplex{Looped{Pulse{X(0.375)},X(0.750)}, Sine{X(1.0 / 400)}})
	Save("EngagedTone.wav",Looped{Multiplex{Compose{Multiplex{Pulse{X(.4)},NewConstant(-6)},Shifted{Pulse{X(.225)},X(.75)}}, Sine{X(1.0/400)}}, X(1.5)})
	Save("RingingTone.wav",Looped{Multiplex{Pulse{X(1)}, Looped{Pulse{X(.4)}, X(.6)}, Stack{Sine{X(1.0/450)},Sine{X(1.0/400)}}}, X(3)})
	Save("NumberUnobtainableTone.wav",Sine{X(1.0/400)})
	Save("dialTone.wav",Stack{Sine{X(1.0/450)},Sine{X(1.0/350)}})

}
/*  hal3 Sat 5 Mar 02:21:24 GMT 2016 go version go1.5.1 linux/386
Sat 5 Mar 02:21:25 GMT 2016 */

