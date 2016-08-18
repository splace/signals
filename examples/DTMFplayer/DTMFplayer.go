// pipe command for converting characters into DTMF tone PCM data.
// example usage:  pipe tones onto aplay;
// ./DTMFplayer[SYSV64].elf <<< "0123456789ABCD#*" | aplay -fs16
// (DTMFplayer has the same default rate as aplay, but has 16bit precision not aplay's 8bit default.)
// or specifiy sample rate;
// ./DTMFplayer\[SYSV64\].elf -rate=16000 <<< "0123456789ABCD#*" | aplay -fs16 -r 16000
package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"bytes"
)

import . "github.com/splace/signals"

var length = X(.07)
var gap = X(.08)

func main() {
	help := flag.Bool("help", false, "display help/usage.")
	var sampleRate uint
	flag.UintVar(&sampleRate, "rate", 8000, "`samples` per second.")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	// cache the raw PCM data for each tone. (helps efficiency if a lot of repeat tones.)  
	var Tones = map[rune]PCM{
		'0': NewPCMSignal(Stacked{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'1': NewPCMSignal(Stacked{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'2': NewPCMSignal(Stacked{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'3': NewPCMSignal(Stacked{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'4': NewPCMSignal(Stacked{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'5': NewPCMSignal(Stacked{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'6': NewPCMSignal(Stacked{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'7': NewPCMSignal(Stacked{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'8': NewPCMSignal(Stacked{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'9': NewPCMSignal(Stacked{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'A': NewPCMSignal(Stacked{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'B': NewPCMSignal(Stacked{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'C': NewPCMSignal(Stacked{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'D': NewPCMSignal(Stacked{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'*': NewPCMSignal(Stacked{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'#': NewPCMSignal(Stacked{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
	}

	var gapPCM = NewPCMSignal(Constant{0}, gap, uint32(sampleRate), 2).(PCM16bit).PCM
	rr := bufio.NewReader(os.Stdin)
	for {
		Rune, _, err := rr.ReadRune()
		if err == io.EOF {
			os.Stdin.Close()
			break
		} else if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, bytes.NewReader(Tones[Rune].Data))
		io.Copy(os.Stdout, bytes.NewReader(gapPCM.Data))
	}
	
	os.Stdout.Close()
}


