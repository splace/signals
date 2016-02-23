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
	Save("BusyTone.wav",Multiplex{Looped{Pulse{375*ms},750*ms}, Sine{UnitX / 400}})
	Save("EngagedTone.wav",Looped{Multiplex{Sum{Multiplex{Pulse{400*ms},NewConstant(-6)},Shifted{Pulse{225*ms},750*ms}}, Sine{UnitX / 400}}, 1500*ms})
	Save("RingingTone.wav",Looped{Multiplex{Pulse{UnitX}, Looped{Pulse{400*ms}, 600*ms}, Stack{Sine{UnitX/450},Sine{UnitX/400}}}, UnitX * 3})
	Save("NumberUnobtainableTone.wav",Sine{UnitX / 400})
	Save("dialTone.wav",Stack{Sine{UnitX/450},Sine{UnitX/350}})

}

