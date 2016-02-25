package signals

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/jpeg" // register de/encoding
	"image/png"  // register de/encoding
	"os"
)

// Images in this package dont have/need a colormodel
// so they are broader, and so can be just cast to, (image.Image isa Image)
// helper types are provided to mix-in colormodel and so implement image.Image.
type Image interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

//var aboveColour,belowColour = color.RGBA{0,0,0,0},color.RGBA{255,255,255,255}
var aboveColour,belowColour = color.RGBA{255,255,255,255},color.RGBA{0,0,0,0}  // upsidedown, as a QAD way to flip image

type FunctionImage struct{
	LimitedFunction
	size image.Rectangle
	yScale y  // cache because cant change for the same size 
}

func (i FunctionImage) Bounds() image.Rectangle{
	return i.size
}

func NewFunctionImage(s LimitedFunction,maxx,maxy int) FunctionImage{
	return FunctionImage{s,image.Rect(0,-maxy,maxx,maxy),Maxy/y(maxy)}
}


func (i FunctionImage) At(xp, yp int) color.Color{
	if i.Call( x(xp) * i.MaxX() / x(i.size.Max.X)-x(i.size.Min.X))<= i.yScale*y(yp)-y(i.size.Min.Y) { 
		return aboveColour
	}
	return belowColour		
}

// wrapper to add colormodel, for Image to conform to image.Image interface
type RGBAImage struct {
	Image
}

func (i RGBAImage) ColorModel() color.Model { return color.RGBAModel }

// wrapper for Image to make it conform to image.Image interface, but allowing down grade to gray, for saving for example.
type GrayImage struct {
	Image
}

func (i GrayImage) ColorModel() color.Model { return color.GrayModel }

// wrapper for Image to make it conform to image.Image interface, but allowing down grade to plan9 paletted, for saving for example.
type Plan9PalettedImage struct {
	Image
}

func (i Plan9PalettedImage) ColorModel() color.Model { return color.Palette(palette.Plan9) }

var Save = savePNG

func savePNG(name string, i Image) {
	wb, err := os.Create(name + ".png")
	if err != nil {
		panic(err)
	}
	defer wb.Close()
	png.Encode(wb, Plan9PalettedImage{i}) 
}

func saveJPG(name string, i Image) {
	wb, err := os.Create(name + ".jpg")
	if err != nil {
		panic(err)
	}
	defer wb.Close()
	jpeg.Encode(wb, RGBAImage{i}, nil)
}


