package signals

import (
	"os"
	"testing"
	"image"
	"image/color"
	"image/color/palette"
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
	png.Encode(file, Plan9PalettedImage{NewDepiction(Multiplex{Sine{UnitX}, Pulse{UnitX}},800,600,color.RGBA{255,255,255,255},color.RGBA{0,0,0,0})})
}
func TestImaging(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	defer stream.Close()
	file, err := os.Create("./test output/M1F1-uint8-AFsp.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
//	png.Encode(wb, Plan9PalettedImage{NewFunctionImage(Pulse{UnitX*4}},3200,300)})   // first second
	jpeg.Encode(file, Plan9PalettedImage{NewDepiction(noise[0],int(noise[0].MaxX()/UnitX*800),600,color.RGBA{255,255,255,255},color.RGBA{0,0,0,0})},nil)    // 800 pixels per second width
}

type composable struct{
	draw.Image
}
func newcomposable(d draw.Image) *composable{
	return &composable{d}
}

func (i *composable) draw(isrc image.Image){
	draw.Draw(i, i.Bounds(),isrc, isrc.Bounds().Min, draw.Src)
} 

func (i *composable) drawAt(isrc image.Image,pt image.Point){
	draw.Draw(i, i.Bounds(),isrc, pt, draw.Src)
} 

func (i *composable) drawOffset(isrc image.Image,pt image.Point){
	draw.Draw(i, i.Bounds(),isrc, isrc.Bounds().Min.Add(pt), draw.Src)
} 

func (i *composable) drawOver(isrc image.Image){
	draw.Draw(i, i.Bounds(),isrc, isrc.Bounds().Min, draw.Over)
} 

func (i *composable) drawOverAt(isrc image.Image,pt image.Point){
	draw.Draw(i, i.Bounds(),isrc, pt, draw.Over)
} 

func (i *composable) drawOverOffset(isrc image.Image,pt image.Point){
	draw.Draw(i, i.Bounds(),isrc, isrc.Bounds().Min.Add(pt), draw.Over)
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
	m := newcomposable(image.NewPaletted(image.Rect(0, -300, 800, 300),palette.WebSafe))
	m.draw(WebSafePalettedImage{NewDepiction(noise[0],800,600,color.RGBA{255,0,0,255},color.RGBA{0,0,0,0})})
	m.drawOver(WebSafePalettedImage{NewDepiction(noise[1],800,600, color.RGBA{0,255,255,127},color.RGBA{0,0,0,0})})
	jpeg.Encode(file, m,nil)
}



