// command for piping from gob encodings of Signals to PCM data.

// example usage (to play a tone):-

// ./player\[SYSV64\].elf < gobs/1kSine.gob | aplay
// or

// cat gobs/1kSine.gob | ./player\[SYSV64\].elf | aplay
//  (1kSine.gob is a procedural 1k cycles sine wave.)

// to specifiy duration:
// ./player\[SYSV64\].elf -length=2 < 1kSine.gob | aplay

// to specifiy sample rate:
// ./player\[SYSV64\].elf -rate=16000 < 1kSine.gob | aplay
// (output s not a higher frequency, since player passes wave format and so includes rate.)

// to specifiy sample precision:
// ./player\[SYSV64\].elf -bytes=1 < 1kSine.gob | aplay
// (bytes can be one of: 1,2,3,4.)
package main

import (
	"bufio"
	"flag"
	"os"
)

import signals "github.com/splace/signals"


func main() {
	help := flag.Bool("help", false, "display help/usage.")
	var sampleRate uint
	flag.UintVar(&sampleRate, "rate", 8000, "`samples` per unit.")
	var samplePrecision uint
	flag.UintVar(&samplePrecision, "bytes", 2, "`bytes` per sample.")
	var length float64
	flag.Float64Var(&length, "length", 1, "length in `units` or Cycles.")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	rr := bufio.NewReader(os.Stdin)
	s,err := signals.Load(rr)
	if err != nil {
		panic("unable to load."+err.Error())
	}
	duration :=signals.X(0)
	switch st:=s.(type) {
	case signals.LimitedSignal:
		if st.MaxX()<=0{
			duration=signals.X(length)
		}else{
			duration=st.MaxX()
		}
	case signals.PeriodicSignal:
		if st.Period()<signals.X(1){
			duration=st.Period()*(signals.MultiplyX(length,signals.X(1))/st.Period())
			}else{
			duration=signals.MultiplyX(length,st.Period())
		}
	case signals.Signal:
			duration=signals.X(length)
	}
	signals.Encode(os.Stdout,uint8(samplePrecision),uint32(sampleRate),duration,s)
	os.Stdout.Close()
}

