// pipe command for converting characters into DTMF tone PCM data.
// example usage:  pipe tones onto aplay;
// ./DTMFplayer[SYSV64].elf <<< "0123456789ABCD#*" | aplay -fs16
// (DTMFplayer has the same default rate as aplay, but has 16bit precision not aplay's 8bit default.)
// or specifiy sample rate;
// ./DTMFplayer\[SYSV64\].elf -rate=16000 <<< "0123456789ABCD#*" | aplay -fs16 -r 16000
// for streaming: "| base64" and "audio/L16"
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
	flag.UintVar(&sampleRate, "rate", 8000, "`samples` per second.")
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
	signals.Encode(os.Stdout,m1,signals.X(1),uint32(sampleRate),2)
	os.Stdout.Close()
}

