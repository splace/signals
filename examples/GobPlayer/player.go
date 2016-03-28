// pipe command for playing gob encodings of functions.
// example usage:  pipe tone to aplay;
// ./player\[SYSV64\].elf < 1kSine.gob | aplay -fs16
//  note  player has the same default rate as aplay, but not the same default precision.
//  note  1kSine.gob is a procedural 1k cycles sine wave.
// or specifiy sample rate:
// ./player\[SYSV64\].elf -rate=16000 < 1kSine.gob | aplay -fs16 -r 16000
// or specifiy duration:
// ./player\[SYSV64\].elf -length=2 < 1kSine.gob | aplay -fs16
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
	var length float64
	flag.Float64Var(&length, "length", 1, "length in `units`")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	m1 := signals.Modulated{}
	rr := bufio.NewReader(os.Stdin)
	if err := m1.Load(rr); err != nil {
		panic("unable to load."+err.Error())
	}
	signals.Encode(os.Stdout,m1,signals.X(length),uint32(sampleRate),2)
	os.Stdout.Close()
}

