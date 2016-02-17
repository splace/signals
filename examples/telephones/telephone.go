package main

import (
	. "../../../signals"
	"os"
)

func Save(file string,s Signal){
	wavFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer wavFile.Close()
	Encode(wavFile,s,UnitTime*4,44100,2)

}

const ms=UnitTime/1000
/*
``On'' and ``off'' durations are in ms. The frequency is 400 Hz, except where noted.
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
	Save("BusyTone.wav",Multiplex{Looped{Pulse{375*ms},750*ms}, Sine{UnitTime / 400}})
	Save("EngagedTone.wav",Looped{Multiplex{Pulse{975*ms}, Looped{Pulse{ 400*ms},750*ms}, Sine{UnitTime / 400}}, 1500*ms})
	Save("RingingTone.wav",Looped{Multiplex{Pulse{UnitTime}, Looped{Pulse{400*ms}, 600*ms}, Stack{Sine{UnitTime/450},Sine{UnitTime/400}}}, UnitTime * 3})
	Save("NumberUnobtainableTone.wav",Sine{UnitTime / 400})
	Save("dialTone.wav",Stack{Sine{UnitTime/450},Sine{UnitTime/350}})

}


