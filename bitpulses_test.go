package signals

import (
	"fmt"
	"math/big"
	"strings"
)

func ExamplePulsePattern() {
	i := new(big.Int)
	_, err := fmt.Sscanf("01110111011101110111011101110111", "%b", i)
	if err != nil {
		panic(i)
	}
	s := PulsePattern{*i, unitX}
	for t := x(0); t < s.MaxX(); t += s.Period(){
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
	 /* Output: 
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
     0.00%                                  X
   100.00%                                                                   X
   100.00%                                                                   X
   100.00%                                                                   X
   */}


