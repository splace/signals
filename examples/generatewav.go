package main

import . "../../signals"
import (
	"fmt"
	"os"
)

func main() {
	m := NewTone(UnitTime/100, 50)
	var file *os.File
	var err error
	if file, err = os.Create(fmt.Sprintf("Sine%+v.wav", m)); err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, 1*UnitTime, 8000, 2)
}

