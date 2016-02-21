package main

import . "../../signals"
import (
	"fmt"
	"os"
)

func main() {
	m := NewTone(UnitTime/100, -6)
	file, err := os.Create(fmt.Sprintf("Sine%+v.wav", m)) // file named after go code of signal
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, 1*UnitTime, 8000, 2)
}

