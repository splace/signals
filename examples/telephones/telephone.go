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
	Save("EngagedTone.wav",Looped{Multiplex{Sum{Multiplex{Pulse{400*ms},NewConstant(50)},Delayed{Pulse{225*ms},750*ms}}, Sine{UnitTime / 400}}, 1500*ms})  //-6db about 50%
	Save("RingingTone.wav",Looped{Multiplex{Pulse{UnitTime}, Looped{Pulse{400*ms}, 600*ms}, Stack{Sine{UnitTime/450},Sine{UnitTime/400}}}, UnitTime * 3})
	Save("NumberUnobtainableTone.wav",Sine{UnitTime / 400})
	Save("dialTone.wav",Stack{Sine{UnitTime/450},Sine{UnitTime/350}})

}

/*  Hal3 Fri Feb 19 00:31:57 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:32:07 GMT 2016 */
/*  Hal3 Fri Feb 19 00:33:45 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:33:50 GMT 2016 */
/*  Hal3 Fri Feb 19 00:35:21 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:35:26 GMT 2016 */
/*  Hal3 Fri Feb 19 00:37:37 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:37:42 GMT 2016 */
/*  Hal3 Fri Feb 19 00:39:38 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:39:43 GMT 2016 */
/*  Hal3 Fri Feb 19 00:42:42 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:42:42 GMT 2016 */
/*  Hal3 Fri Feb 19 00:42:58 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:42:59 GMT 2016 */
/*  Hal3 Fri Feb 19 00:43:45 GMT 2016 go version go1.5.1 linux/amd64
Fri Feb 19 00:43:50 GMT 2016 */

