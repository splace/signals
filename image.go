package signals

import (
	"image"
	"image/color"
	"image/color/palette"
)

// Depiction of a LimitedFunction   
type Depiction struct{
	LimitedFunction
	size image.Rectangle
	belowColour,aboveColour  color.Color
	yScale y  // optimisation; small, a cache because this can't change for the same size, which is not exposed. 
}

func (i Depiction) Bounds() image.Rectangle{
	return i.size
}

// makes an image of a LimitedSignal, scaled to maxx x maxy pixels.
func NewDepiction(s LimitedFunction,maxx,maxy int, c1,c2 color.Color) Depiction{
	return Depiction{s,image.Rect(0,-maxy/2,maxx,maxy/2), c1,c2,Maxy/y(maxy/2)}
}

func (i Depiction) At(xp, yp int) color.Color{
	if i.Call( x(xp) * i.MaxX() / x(i.size.Max.X)-x(i.size.Min.X))<= i.yScale*y(yp)-y(i.size.Min.Y) { 
		return i.aboveColour
	}
	return i.belowColour		
}


// an depiction is an image.Image without a colormodel, so is more general.
// embedded in helper wrapper to implement image.Image.
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

