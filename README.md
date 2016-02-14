# signals

overview:

https://github.com/splace/signals/blob/master/doc.go	

(included in doc below)

installation:

     go get github.com/splace/signals   

Example:

	package main
	import "github.com/splace/signals"
	import ("fmt","os")
	
	func main() {
		m := signals.NewTone(UnitTime/100, 50)
		var file *os.File
		var err error
		if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
			panic(err)
		}
		defer file.Close()
		Encode(file, m, UnitTime, 8000, 1)
	}

output:
	"https://github.com/splace/signals/blob/master/Sine[{Cycle:%20%20%20%20%200.01s}%20{Constant:%20%20%20%2050.00%}].wav"

status:

Signal generators:- Sine,Square,Pulse,Heavyside,Bittrain,ADSR,Constant,RampUp,RampDown,Sigmoid

Signal modifiers:- Delay,Spedup,Looped,Inverted,Reversed,Modulated,TriggerRising

Signal Combiners:- Add,Multi

docs: 
     
[![GoDoc](https://godoc.org/github.com/splace/signals?status.svg)](https://godoc.org/github.com/splace/signals)

