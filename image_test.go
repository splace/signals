package signals

import (
	"os"
	"testing"
	"image/color"
)

func TestImagingSine(t *testing.T) {
	Save("./test output/Sine",NewFunctionImage(Multiplex{Sine{UnitX}, Pulse{UnitX}},800,600))
}
func TestImaging(t *testing.T) {
	stream, err := os.Open("M1F1-uint8-AFsp.wav")
	if err != nil {
		panic(err)
	}
	noise, err := Decode(stream)
	aboveColour = color.RGBA{0,0,255,255}

	defer stream.Close()
//	Save("./test output/M1F1-uint8-AFsp.wav",NewFunctionImage(Multiplex{noise[0], Pulse{UnitX*4}},3200,300))   // first second
	Save("./test output/M1F1-uint8-AFsp.wav",NewFunctionImage(noise[0],int(noise[0].MaxX()/UnitX*800),600))    // 800 pixels per second width
}


