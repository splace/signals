# Signals

Status: (Beta :- stabilising API)


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
	m := Modulated{Sine{OneSecond/100},NewConstant(-6)}
	// save file named after the go code of the signal
	file, err := os.Create(fmt.Sprintf("%+v.wav", m)) 
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, 2, 8000, OneSecond, m)
}
```
Output: Sine wave, 100hz, 50% volume (-6dB), 1 sec, @8k samples/sec, 2byte signed PCM (s16), WAV file 

[[{Cycle:     0.01s} {Constant:    50.00%}].wav](https://github.com/splace/signals/blob/master/examples/%5B%7BCycle:%20%20%20%20%200.01s%7D%20%7BConstant:%20%20%20%2050.00%25%7D%5D.wav)

Features:

  * sources:- Sine,Square,Pulse,Heavyside,Bittrain,RampUp,RampDown,Sigmoid,PCM{bits}bit
(PCM sources can be loaded from wav files)
	
  * modifiers:- Delayed,Spedup,Looped,Inverted,Reversed,RateModulated,Triggered,Segmented

  * combiners:- Stack,Composite,Modulate

  * extras(non-core):- Depiction,ADSR,Noise


Extras examples: Depiction of Stack{Sine{unitX/100}, Sine{unitX/50}}, red/black,(3200px.600px) for 4 * unitX. 

(see ![image tests](https://github.com/splace/signals/blob/master/image_test.go) )

![speech saved as wav](https://github.com/splace/signals/blob/master/test%20output/out.jpeg)
