# signals

Overview: (see godoc reference below)

Installation:

     go get github.com/splace/signals   

Example:
```go
package main
import (
	"fmt"
	"os"
)

import . "github.com/splace/signals"

func main() {
	m := NewTone(UnitTime/100, 50)
	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, 1*UnitTime, 8000, 1)
}
```
Output: 1 sec, 100hz, 50% volume, Sine wave, @8k samples/sec, 8bit unsigned PCM (u8), WAV file 

[Sine[{Cycle:     0.01s} {Constant:    50.00%}].wav](https://github.com/splace/signals/blob/master/examples/Sine%5B%7BCycle:%20%20%20%20%200.01s%7D%20%7BConstant:%20%20%20%2050.00%25%7D%5D.wav)

Status:

generators:- Sine,Square,Pulse,Heavyside,Bittrain,ADSR,RampUp,RampDown,Sigmoid,Noise

modifiers:- Delayed,Spedup,Looped,Inverted,Reversed,Modulated,Triggered,Segmented

combiners:- Stack,Multiplex,Sum

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/signals?status.svg)](https://godoc.org/github.com/splace/signals)

