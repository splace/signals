package main

import . "github.com/splace/signals"  //  "../../signals" //
import (
	"fmt"
	"os"
)

func main() {
	m := NewTone(X(1.0/100), -6)
	file, err := os.Create(fmt.Sprintf("Sine%+v.wav", m)) // file named after go code of generating Function
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Encode(file, m, X(1), 8000, 2)
}

