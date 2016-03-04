package signals

import (
	"image"
	"image/color"
	"image/color/palette"
)

// a Depictor is an image.Image without a colormodel, so is more general.
// embedded in one of the helper wrappers to get an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// Depiction is a Depictor of a Function
type Depiction struct {
	Function
	size                     image.Rectangle
	pixelsPerUnitX           int
	belowColour, aboveColour color.Color
}

// makes a Depiction of a LimitedSignal, scaled to pxMaxx x pxMaxy pixels and sets its colours.
func NewDepiction(s LimitedFunction, pxMaxX, pxMaxY int, c1, c2 color.Color) Depiction {
	return Depiction{s, image.Rect(0, -pxMaxY/2, pxMaxX, pxMaxY/2), int(int64(pxMaxX) * int64(unitX) / int64(s.MaxX())), c1, c2}
}

func (i Depiction) Bounds() image.Rectangle {
	return i.size
}

func (i Depiction) At(xp, yp int) color.Color {
	if i.call(x(xp)*unitX/x(i.pixelsPerUnitX)-x(i.size.Min.X)) <= maxY/y(i.size.Max.Y)*y(yp)-y(i.size.Min.Y) {
		return i.aboveColour
	}
	return i.belowColour
}

// RGBA depiction wrapper
type RGBAImage struct {
	Depictor
}

func (i RGBAImage) ColorModel() color.Model { return color.RGBAModel }

// gray depiction wrapper.
type GrayImage struct {
	Depictor
}

func (i GrayImage) ColorModel() color.Model { return color.GrayModel }

// plan9 paletted, depiction wrapper.
type Plan9PalettedImage struct {
	Depictor
}

func (i Plan9PalettedImage) ColorModel() color.Model { return color.Palette(palette.Plan9) }

// WebSafe paletted, depiction wrapper.
type WebSafePalettedImage struct {
	Depictor
}

func (i WebSafePalettedImage) ColorModel() color.Model { return color.Palette(palette.WebSafe) }
