// convert a stereo wav file into a mono by adding sounds together.
/*
Usage :
 -bytes precision
    	precision in bytes per sample. (requires format option set) (default 2)
  -chans string
    	extract/recombine listed channel number(s) only. (default "0,1")
  -db uint
    	reduce volume in dB (6 to halve.) stacked channels could clip without.
  -format
    	don't use input sample rate and precision for output, use command-line options
  -help
    	display help/usage.
  -prefix string
    	add individual prefixes to extracted mono file(s) names. (default "L-,R-")
  -rate samples
    	samples per second.(requires format option set) (default 44100)
  -stack
    	recombine all channels into a mono file.
*/
package main

import . "github.com/splace/signals" //"../../../signals"  //
import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

// Note: experiment with a fancy bespoke logger

type messageLog struct {
	*log.Logger
	message string
}

// pick off and handle the returned error (second arg.) to simplify handling structure.
func (ml messageLog) errFatal(result interface{}, err error) interface{} {
	if err != nil {
		ml.Fatal(err.Error())
	}
	return result
}

func (ml messageLog) Fatal(info string) {
	ml.Logger.Fatal("\t" + os.Args[0] + "\t" + ml.message + "\t" + info)
	return
}

/*
DEBUG1..DEBUG5 	Provides successively-more-detailed information for use by developers.
INFO 	Provides information implicitly requested by the user, e.g., output from VACUUM VERBOSE.
NOTICE 	Provides information that might be helpful to users, e.g., notice of truncation of long identifiers.
WARNING 	Provides warnings of likely problems, e.g., COMMIT outside a transaction block.
ERROR 	Reports an error that caused the current command to abort.
LOG 	Reports information of interest to administrators, e.g., checkpoint activity.
FATAL 	Reports an error that caused the current session to abort.
PANIC 	Reports an error that caused all database sessions to abort.
*/

func main() {
	format := flag.Bool("format", false, "don't use input sample rate and precision for output, use flag(s)")
	stack := flag.Bool("stack", false, "recombine all channels into a mono file.")
	help := flag.Bool("help", false, "display help/usage.")
	var dB uint
	flag.UintVar(&dB, "db", 0, "reduce volume in dB (6 to halve.) stacked channels could clip without.")
	var channels, namePrefix string
	flag.StringVar(&channels, "chans", "1,2", "extract/recombine listed channel number(s) only. ('1,2' for first 2 channels)")
	flag.StringVar(&namePrefix, "prefix", "L-,R-,C-,LFE-,LB-,RB-", "add individual prefixes to extracted mono file(s) names.")
	var sampleRate, sampleBytes uint
	flag.UintVar(&sampleRate, "rate", 44100, "`samples` per second.(requires format option set)")
	flag.UintVar(&sampleBytes, "bytes", 2, "`precision` in bytes per sample. (requires format option set)")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	files := flag.Args()
	errorLog := messageLog{log.New(os.Stderr, "ERROR\t", log.LstdFlags), "File access"} // 'ERROR' log for, initially, file access.
	var in, out *os.File
	if len(files) == 2 {
		in = errorLog.errFatal(os.Open(files[0])).(*os.File)
		defer in.Close()
	} else {
		errorLog.Fatal("2 file names required.")
	}
	errorLog.message = "Decode:" + files[0]
	PCMs := errorLog.errFatal(Decode(in)).([]PeriodicLimitedSignal)
	if *format {
		if *stack {
			errorLog.message = "File Access"
			out = errorLog.errFatal(os.Create(files[1])).(*os.File)
			errorLog.message = "Encode"
			Encode(out, uint8(sampleBytes), uint32(sampleRate), PCMs[0].MaxX(), Modulated{NewStack(PromoteToSignals(PCMs)...), NewConstant(float64(-dB))})
			out.Close()
		} else {
			errorLog.message = "Parse Channels."
			chs := map[int]struct{}{}
			for _, c := range strings.Split(channels, ",") {
				chs[int(errorLog.errFatal(strconv.ParseUint(c, 10, 16)).(uint64))] = struct{}{}
			}
			prefixes := strings.Split(namePrefix, ",")
			for i, n := range PCMs {
				if _, ok := chs[i]; !ok {
					continue
				}
				errorLog.message = "File Access"
				out = errorLog.errFatal(os.Create(prefixes[i] + files[1])).(*os.File)
				errorLog.message = "Encode"
				Encode(out, uint8(sampleBytes), uint32(sampleRate), n.MaxX(), Modulated{n, NewConstant(float64(-dB))})
				out.Close()
			}
		}
	} else {
		if *stack {
			errorLog.message = "File Access"
			out = errorLog.errFatal(os.Create(files[1])).(*os.File)
			errorLog.message = "Encode"
			EncodeLike(out, Modulated{NewStack(PromoteToSignals(PCMs)...), NewConstant(float64(-dB))}, PCMs[0])
			out.Close()
		} else {
			errorLog.message = "Parse Channels."
			chs := map[int]struct{}{}
			for _, c := range strings.Split(channels, ",") {
				chs[int(errorLog.errFatal(strconv.ParseUint(c, 10, 16)).(uint64))-1] = struct{}{}
			}
			prefixes := strings.Split(namePrefix, ",")
			for i, n := range PCMs {
				if _, ok := chs[i]; !ok {
					continue
				}
				errorLog.message = "File Access"
				out = errorLog.errFatal(os.Create(prefixes[i] + files[1])).(*os.File)
				errorLog.message = "Encode"
				EncodeLike(out, Modulated{n, NewConstant(float64(-dB))}, PCMs[0])
				out.Close()
			}
		}
	}
}



