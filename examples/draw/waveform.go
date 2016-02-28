// use of mdlayher/waveform
package main

import "github.com/mdlayher/waveform"
import . "github.com/splace/signals"
import (
	"fmt"
	"os"
	"io"
	"image/color"
	"image/png"  // register de/encoding

)

func main() {
	m := NewTone(UnitX/50, 0)

	out, in := io.Pipe()
	go func() {
		Encode(in, m, UnitX/25,22100, 1)
		in.Close()
	}()
	
	bColor:=waveform.BGColorFunction(waveform.SolidColor(color.Black))
	fColor:=waveform.FGColorFunction(waveform.SolidColor(color.White))
	img, err := waveform.Generate(out,waveform.Resolution(22100),bColor,fColor)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
    	fmt.Println(err)
	}
}



