// make a jpeg image from a stereo wav file.
// usage: 2jpeg.<<o|exe>> <<stereo.wav>>
package main

import . "github.com/splace/signals"  //"../../../signals" // 
import (
	"flag"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/jpeg" // register de/encoding
	"log"
	"os"
)

func main() {
	//var sampleRate,sampleBytes uint
	//flag.UintVar(&sampleRate, "rate", 44100, "sample per second")
	//flag.UintVar(&sampleBytes,"bytes", 2, "bytes per sample")
	flag.Parse()
	files := flag.Args()
	myLog := messageLog{log.New(os.Stderr, "ERROR\t", log.LstdFlags), "File access"}
	var in, out *os.File
	if len(files) != 1 {
		myLog.Fatal("1 file names required.")
	}
	in=myLog.errFatal(os.Open(files[0])).(*os.File)
	myLog.message = "Decode:" + files[0]
	defer in.Close()
	noise := myLog.errFatal(Decode(in)).([]LimitedFunction)
	if len(noise) != 2 {
		myLog.Fatal("Need a stereo input file.")
	}
	myLog.message = "File Access"
	out = myLog.errFatal(os.Create(files[0]+".jpeg")).(*os.File)
	defer out.Close()
	m := newcomposable(image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe))
	// offset centre of 600px image, to fit 300px width.
	m.drawOffset(WebSafePalettedImage{NewDepiction(noise[0], 800, 600, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	m.drawOverOffset(WebSafePalettedImage{NewDepiction(noise[1], 800, 600, color.RGBA{0, 255, 255, 127}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	jpeg.Encode(out, m, nil)
}

type messageLog struct {
	*log.Logger
	message string
}

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

type composable struct {
	draw.Image
}

func newcomposable(d draw.Image) *composable {
	return &composable{d}
}

func (i *composable) draw(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Src)
}

func (i *composable) drawAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Src)
}

func (i *composable) drawOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Src)
}

func (i *composable) drawOver(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Over)
}

func (i *composable) drawOverAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Over)
}

func (i *composable) drawOverOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Over)
}

