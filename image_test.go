package signals

import (
	"os"
	"testing"
	"image"
	"image/color"
	"image/jpeg" // register de/encoding
	"image/png"  // register de/encoding
	"image/draw"
)

func TestImagingSine(t *testing.T) {
	file, err := os.Create("./test output/Sine.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, Plan9PalettedImage{NewDepiction(Multiplex{Sine{UnitX}, Pulse{UnitX}},800,600)})
}
func TestImaging(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	aboveColour = color.RGBA{0,0,255,255}

	defer stream.Close()
	file, err := os.Create("./test output/M1F1-uint8-AFsp.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
//	png.Encode(wb, Plan9PalettedImage{NewFunctionImage(Pulse{UnitX*4}},3200,300)})   // first second
	jpeg.Encode(file, Plan9PalettedImage{NewDepiction(noise[0],int(noise[0].MaxX()/UnitX*800),600)},nil)    // 800 pixels per second width
}
func TestI(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	defer stream.Close()
	file, err := os.Create("./test output/out.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	m := image.NewRGBA(image.Rect(0, -300, 800, 300))
	aboveColour = color.RGBA{0,0,0,0}
	belowColour = color.RGBA{0,0,255,255}
	src:= Plan9PalettedImage{NewDepiction(noise[0],800,600)}
	draw.Draw(m, m.Bounds(),src, src.Bounds().Min, draw.Over)
	aboveColour = color.RGBA{0,0,0,0}
	belowColour = color.RGBA{255,0,0,255}
	src2:= Plan9PalettedImage{NewDepiction(noise[1],800,600)}
	draw.Draw(m, m.Bounds(),src2, image.Pt(src2.Bounds().Min.X,src2.Bounds().Min.Y+50), draw.Over)
	jpeg.Encode(file, m,nil)
}



