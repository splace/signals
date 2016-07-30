package signals

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/jpeg" // register de/encoding
	"image/png"  // register de/encoding
	"os"
	"testing"
)

func TestImageSine(t *testing.T) {
	file, err := os.Create("./test output/Sine.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, Plan9PalettedImage{NewDepiction(Modulated{Sine{unitX}, Pulse{unitX}}, 800, 600, color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 0})})
}

func TestImage(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	defer stream.Close()
	file, err := os.Create("./test output/outp.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//	png.Encode(wb, Plan9PalettedImage{NewFunctionImage(Pulse{UnitX*4}},3200,300)})   // first second
	jpeg.Encode(file, Plan9PalettedImage{Depiction{noise[0], image.Rect(0, -300, int(noise[0].MaxX()*800/unitX), 300), 800, color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 0}}}, nil) // 800 pixels per second width
}

func TestImageComposable(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	defer stream.Close()
	file, err := os.Create("./test output/outsp.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	m := &composable{image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe)}
	// offset centre of 600px image, to fit 300px width.
	m.drawOffset(WebSafePalettedImage{NewDepiction(noise[0], 800, 600, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	m.drawOverOffset(WebSafePalettedImage{NewDepiction(noise[1], 800, 600, color.RGBA{0, 255, 255, 127}, color.RGBA{0, 0, 0, 0})}, image.Point{0, 150})
	jpeg.Encode(file, m, nil)
}

func TestImageStack(t *testing.T) {
	s := Stacked{Sine{unitX / 100}, Sine{unitX / 50}}
	file, err := os.Create("./test output/out.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ds := Depiction{s, image.Rect(0, -150, 800, 150), 3200, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0}}

	m := &composable{image.NewPaletted(ds.Bounds(), palette.WebSafe)}
	m.draw(WebSafePalettedImage{ds})
	jpeg.Encode(file, m, nil)
}

func TestImageMultiplex(t *testing.T) {
	s := Modulated{Sine{unitX / 100}, Sine{unitX / 50}}
	file, err := os.Create("./test output/multiplex.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ds := Depiction{s, image.Rect(0, -150, 800, 150), 3200, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 0, 0}}

	m := &composable{image.NewPaletted(ds.Bounds(), palette.WebSafe)}
	m.draw(WebSafePalettedImage{ds})
	jpeg.Encode(file, m, nil)
}

// composable is a draw.Image that comes with helper functions to simplify Draw function.
type composable struct {
	draw.Image
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


