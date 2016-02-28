// convert a stereo wav file into a mono by adding sounds together.
// usage: 2mono.<<o|exe>> <<stereo.wav>> <<mono.wav>>
// doesn't need anything that is 'sound' specific, just treats as abstract PCM data.
// Note: experiment with a fancy bespoke logger
package main

import . "../../../signals"  // github.com/splace/signals //
import (
	"os"
	"flag"
	"log"
)


type messageLog struct{
	*log.Logger
	message string
}


func (ml messageLog) errFatal(result interface{},err error) interface{}{
	if err!=nil{
		ml.Fatal(err.Error())
	}
	return result
}

func (ml messageLog) Fatal(info string) {
	ml.Logger.Fatal("\t"+os.Args[0]+"\t"+ml.message+"\t"+info)
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
	//var sampleRate,sampleBytes uint
	//flag.UintVar(&sampleRate, "rate", 44100, "sample per second")
	//flag.UintVar(&sampleBytes,"bytes", 2, "bytes per sample")
	flag.Parse()
	files := flag.Args()
	myLog := messageLog{log.New(os.Stderr,"ERROR\t",log.LstdFlags),"File access"} 
	var in,out *os.File
	if len(files)==2 {
		in=myLog.errFatal(os.Open(files[0])).(*os.File)
		defer in.Close()
	}else{
		myLog.Fatal( "2 file names required.")
	}
	myLog.message="Decode:"+files[0]
	noise:=myLog.errFatal(Decode(in)).([]Function)
	if len(noise)!=2{
		myLog.Fatal("Need a stereo input file.")
	}
	myLog.message="File Access"
	out=myLog.errFatal(os.Create(files[1])).(*os.File)
	defer out.Close()
	// save stacked channels with the same sample Rate and precision as the first channel
	EncodeLike(out,Stack{noise[0],noise[1]},noise[0].(PCMFunction))		
}


