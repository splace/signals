package main

import . "../../signals"
import (
	"fmt"
	"os"
)

func main() {
	m := NewTone(UnitX/100, -6)
	file, err := os.Create(fmt.Sprintf("Sine%+v.wav", m)) // file named after go code of generating Function
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, 1*UnitX, 8000, 2)
}

