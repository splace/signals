package signals

import (
	"image"
	"image/color"
	"image/color/palette"
)

// Depiction is a image.Image of a LimitedFunction   
type Depiction struct{
	LimitedFunction
	size image.Rectangle
	belowColour,aboveColour  color.Color
	yScale y  //small optimisation;  a cache because this can't change for the same size, which is not exposed. 
}
// makes an image of a LimitedSignal, scaled to maxx x maxy pixels.
func NewDepiction(s LimitedFunction,maxx,maxy int, c1,c2 color.Color) Depiction{
	return Depiction{s,image.Rect(0,-maxy/2,maxx,maxy/2), c1,c2,Maxy/y(maxy/2)}
}

func (i Depiction) Bounds() image.Rectangle{
	return i.size
}


func (i Depiction) At(xp, yp int) color.Color{
	if i.Call( x(xp) * i.MaxX() / x(i.size.Max.X)-x(i.size.Min.X))<= i.yScale*y(yp)-y(i.size.Min.Y) { 
		return i.aboveColour
	}
	return i.belowColour		
}


// a depiction is an image.Image without a colormodel, so is more general.
// embedded in one of the helper wrappers to get an image.Image.
type depiction interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// RGBA depiction wrapper
type RGBAImage struct {
	depiction
}

func (i RGBAImage) ColorModel() color.Model { return color.RGBAModel }

// gray depiction wrapper.
type GrayImage struct {
	depiction
}

func (i GrayImage) ColorModel() color.Model { return color.GrayModel }

// plan9 paletted, depiction wrapper.
type Plan9PalettedImage struct {
	depiction
}

func (i Plan9PalettedImage) ColorModel() color.Model { return color.Palette(palette.Plan9) }
// WebSafe paletted, depiction wrapper.
type WebSafePalettedImage struct {
	depiction
}

func (i WebSafePalettedImage) ColorModel() color.Model { return color.Palette(palette.WebSafe) }

