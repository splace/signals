package signals

import (
	"fmt"
	"strings"
)

func PrintGraph2(s Signal, start, end, step x) {
	for t := start; t < end; t += step {
		fmt.Println(s.property(t), strings.Repeat(" ", int(s.property(t)/(unitY/33))+33)+"X")
	}
}

func ExampleCombinersSequenced() {
	PrintGraph2(NewSequence(Pulse{unitX},Pulse{unitX}), 0, 3*unitX, unitX/4)
	/* Output:
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/

}
