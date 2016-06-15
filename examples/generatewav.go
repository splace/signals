package main

import . "github.com/splace/signals"
import (
	"fmt"
	"os"
)

var OneSecond = X(1)

func main() {
	signal := Modulated{Sine{OneSecond/100},NewConstant(-6)}
	// save file named after the go code of the signal
	file, err := os.Create(fmt.Sprintf("%+v.wav", signal)) 
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, 2, 8000, OneSecond, signal)
}

