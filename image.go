package signals

import (
	"image"
	"image/color"
	"image/color/palette"
)


var aboveColour,belowColour = color.RGBA{0,0,0,0},color.RGBA{255,255,255,255}

// Depiction of a LimitedFunction   
type Depiction struct{
	LimitedFunction
	size image.Rectangle
	yScale y  // optimisation; small, a cache because this can't change for the same size, which is not exposed. 
}

func (i Depiction) Bounds() image.Rectangle{
	return i.size
}

// makes an image of a LimitedSignal, scaled to maxx x maxy pixels.
func NewDepiction(s LimitedFunction,maxx,maxy int) Depiction{
	return Depiction{s,image.Rect(0,-maxy/2,maxx,maxy/2),Maxy/y(maxy/2)}
}

func (i Depiction) At(xp, yp int) color.Color{
	if i.Call( x(xp) * i.MaxX() / x(i.size.Max.X)-x(i.size.Min.X))<= i.yScale*y(yp)-y(i.size.Min.Y) { 
		return aboveColour
	}
	return belowColour		
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
/*  Hal3 Fri Feb 26 21:39:58 GMT 2016 go version go1.5.1 linux/amd64
FAIL	_/home/simon/Dropbox/github/working/signals [build failed]
Fri Feb 26 21:39:59 GMT 2016 */

