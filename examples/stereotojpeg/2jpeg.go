// make a jpeg image from a stereo wav file.
// usage: 2jpeg.<<elf|exe>> <<stereo.wav>>
package main

import . "github.com/splace/signals" 
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
	flag.Parse()
	files := flag.Args()

	log.SetPrefix("ERROR\tFile access.\t")
	var in, out *os.File
	AssertFatal(func ()bool{return len(files)==1}," to detect an input file name.")
	in = ErrFatal(os.Open(files[0])).(*os.File)
	defer in.Close()

	log.SetPrefix("ERROR\tDecode:"+ files[0]+"\t")
	noise := ErrFatal(Decode(in)).([]PeriodicLimitedSignal)
	AssertFatal(func ()bool{return len(noise) ==2}," to detect a stereo file.")

	log.SetPrefix("ERROR\tFile Access.\t")
	out = ErrFatal(os.Create(files[0] + ".jpeg")).(*os.File)
	defer out.Close()

	m := newcomposable(image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe))
	// offset centre of 600px image, to fit 300px width.
	m.drawOffset(WebSafePalettedImage{NewDepiction(noise[0], 800, 600, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	m.drawOverOffset(WebSafePalettedImage{NewDepiction(noise[1], 800, 600, color.RGBA{0, 255, 255, 127}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	jpeg.Encode(out, m, nil)
}


//helpers for log

func AssertFatal(test func()bool,info string) {
	if !test() {
		log.Fatal("failed"+info)
	}
	return
}

func ErrFatal(result interface{}, err error) interface{} {
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
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

// helpers for drawing

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


