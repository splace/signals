# functions

Status: (Beta :- stabilising API)

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
	m := NewTone(UnitX/100, -6)
	file, err := os.Create(fmt.Sprintf("Sine%+v.wav", m)) // file named after go code of generating Function
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, 1*UnitX, 8000, 2)
}

```
Output: Sine wave, 100hz, 50% volume (-6dB), 1 sec, @8k samples/sec, 2byte signed PCM (s16), WAV file 

[Sine[{Cycle:     0.01s} {Constant:    50.00%}].wav](https://github.com/splace/signals/blob/master/examples/Sine%5B%7BCycle:%20%20%20%20%200.01s%7D%20%7BConstant:%20%20%20%2050.00%25%7D%5D.wav)

Features:

generators:- Sine,Square,Pulse,Heavyside,Bittrain,ADSR,RampUp,RampDown,Sigmoid,Noise,PCM<<bits>>bit

modifiers:- Delayed,Spedup,Looped,Inverted,Reversed,Modulated,Triggered,Segmented

combiners:- Stack,Sum,Multiplex

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/signals?status.svg)](https://godoc.org/github.com/splace/signals)

