package signals

import (
	"os"
	"testing"
	"image/color"
	"image/jpeg" // register de/encoding
	"image/png"  // register de/encoding
)

func TestImagingSine(t *testing.T) {
	file, err := os.Create("./test output/Sine.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, Plan9PalettedImage{NewFunctionImage(Multiplex{Sine{UnitX}, Pulse{UnitX}},800,600)})
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
	jpeg.Encode(file, Plan9PalettedImage{NewFunctionImage(noise[0],int(noise[0].MaxX()/UnitX*800),600)},nil)    // 800 pixels per second width
}

func TestStereoImaging(t *testing.T) {
	stream, err := os.Open("drmapan.wav")
	if err != nil {
		panic(err)
	}
	channels, err := Decode(stream)
	aboveColour = color.RGBA{0,0,255,255}

	defer stream.Close()
	file, err := os.Create("./test output/drmapan.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
//	png.Encode(wb, Plan9PalettedImage{NewFunctionImage(Pulse{UnitX*4}},3200,300)})   // first second
	jpeg.Encode(file, Plan9PalettedImage{NewFunctionImage(channels[0],int(channels[0].MaxX()/UnitX*800),600)},nil)    // 800 pixels per second width
}


