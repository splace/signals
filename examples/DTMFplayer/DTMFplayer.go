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
	// cache the raw PCM data for each tone
	// (example, not really required.)
	var Tones = map[rune]PCM{
		'0': NewPCMFunction(Stack{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'1': NewPCMFunction(Stack{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'2': NewPCMFunction(Stack{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'3': NewPCMFunction(Stack{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'4': NewPCMFunction(Stack{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'5': NewPCMFunction(Stack{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'6': NewPCMFunction(Stack{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'7': NewPCMFunction(Stack{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'8': NewPCMFunction(Stack{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1336)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'9': NewPCMFunction(Stack{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'A': NewPCMFunction(Stack{Sine{X(1.0 / 697)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'B': NewPCMFunction(Stack{Sine{X(1.0 / 770)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'C': NewPCMFunction(Stack{Sine{X(1.0 / 852)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'D': NewPCMFunction(Stack{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1633)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'*': NewPCMFunction(Stack{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1209)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
		'#': NewPCMFunction(Stack{Sine{X(1.0 / 941)}, Sine{X(1.0 / 1477)}}, length, uint32(sampleRate), 2).(PCM16bit).PCM,
	}

	var gapPCM = NewPCMFunction(Constant{0}, gap, uint32(sampleRate), 2).(PCM16bit).PCM
	rr := bufio.NewReader(os.Stdin)
	for {
		Rune, _, err := rr.ReadRune()
		if err == io.EOF {
			os.Stdin.Close()
			break
		} else if err != nil {
			panic(err)
		}
		Tones[Rune].Encode(os.Stdout)
		gapPCM.Encode(os.Stdout)
	}
	os.Stdout.Close()
}


