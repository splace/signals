// make a jpeg image from a stereo wav file.
// usage: 2jpeg.<<elf|exe>> <<stereo.wav>>
package main

import . "github.com/splace/signals" //"../../../signals" //
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
	sLogError := statefulLogger{log.New(os.Stderr, "ERROR\t", log.LstdFlags), "File access"}
	var in, out *os.File
	sLogError.AssertFatal(" to detect an input file name.",func ()bool{return len(files)==1})
	in = sLogError.ErrFatal(os.Open(files[0])).(*os.File)
	sLogError.State = "Decode:" + files[0]
	defer in.Close()
	noise := sLogError.ErrFatal(Decode(in)).([]PCMSignal)
	sLogError.AssertFatal(", input file not stereo.",func ()bool{return len(noise) ==2})
	sLogError.State = "File Access"
	out = sLogError.ErrFatal(os.Create(files[0] + ".jpeg")).(*os.File)
	defer out.Close()
	m := newcomposable(image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe))
	// offset centre of 600px image, to fit 300px width.
	m.drawOffset(WebSafePalettedImage{NewDepiction(noise[0], 800, 600, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	m.drawOverOffset(WebSafePalettedImage{NewDepiction(noise[1], 800, 600, color.RGBA{0, 255, 255, 127}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	jpeg.Encode(out, m, nil)
}

type statefulLogger struct {
	*log.Logger
	State string
}

func (sl statefulLogger) ErrFatal(result interface{}, err error) interface{} {
	if err != nil {
		sl.Fatal(err.Error())
	}
	return result
}

func (sl statefulLogger) AssertFatal(info string,test func()bool) {
	if !test() {
		sl.Fatal("failed"+info)
	}
	return
}

func (sl statefulLogger) Fatal(info string) {
	sl.Logger.Fatal("\t" + os.Args[0] + "\t" + sl.State + "\t" + info)
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


