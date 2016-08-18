# Signals

Status: (Beta :- stabilising API)

see "examples" folder for some uses.

Overview/docs: [![GoDoc](https://godoc.org/github.com/splace/signals?status.svg)](https://godoc.org/github.com/splace/signals) 

Installation:

     go get github.com/splace/signals   

Example:
```go
package main

import . "github.com/splace/signals"
import (
	"fmt"
	"os"
)

var OneSecond = X(1)

func main() {
	signal := Modulated{Sine{OneSecond/100},NewConstant(-6)}
	// save file named after the go code of the signal
	file, err := os.Create(fmt.Sprintf("%+v.wav", signal)) 
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, 2, 8000, OneSecond, signal)
}
```
Output: Sine wave, 100hz, 50% volume (-6dB), 1 sec, @8k samples/sec, 2byte signed PCM (s16), WAV file 

Features:

  * sources:- Sine, Square, Pulse, Heavyside, Bittrain, RampUp, RampDown, Sigmoid, PCM{8|16|24|32|48}bit (PCM sources can be stored in wav files)
	
  * modifiers:- Delayed, Spedup, Looped, Inverted, Reversed, Cached, RateModulated, Triggered, Segmented

  * combiners:- Sequenced, Modulated, Stacked, Composite

  * extras(non-core):- Depiction, ADSR, Noise, Wave (stream)


Extras examples: Depiction of "Stack{Sine{unitX/100}, Sine{unitX/50}}", red/black,(3200px.600px) for 4 * unitX. 

(see ![image tests](https://github.com/splace/signals/blob/master/image_test.go) )

![speech saved as wav](https://github.com/splace/signals/blob/master/test%20output/out.jpeg)
