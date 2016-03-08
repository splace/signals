package main

import . "github.com/splace/signals"
import (
	"fmt"
	"os"
)

var OneSecond = X(1)

func main() {
	m := NewTone(OneSecond/100, -6)
	file, err := os.Create(fmt.Sprintf("Sine%+v.wav", m)) // file named after go code of generating Function
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, OneSecond, 8000, 2)
}

