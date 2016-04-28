package signals

import (
	"fmt"
	"strings"
)

func ExampleADSREnvelope() {
	s := NewADSREnvelope(unitX, unitX, unitX, unitY/2, unitX)
	for t := x(0); t < 5*unitX; t += unitX / 10 {
		fmt.Println(s.property(t), strings.Repeat(" ", int(s.property(t)/(unitY/33))+33)+"X")
	}
	fmt.Println()
	/* Output:
   0.00%                                  X
  10.00%                                     X
  20.00%                                        X
  30.00%                                           X
  40.00%                                               X
  50.00%                                                  X
  60.00%                                                     X
  70.00%                                                         X
  80.00%                                                            X
  90.00%                                                               X
 100.00%                                                                  X
  95.00%                                                                 X
  90.00%                                                               X
  85.00%                                                              X
  80.00%                                                            X
  75.00%                                                          X
  70.00%                                                         X
  65.00%                                                       X
  60.00%                                                     X
  55.00%                                                    X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  50.00%                                                  X
  45.00%                                                X
  40.00%                                               X
  35.00%                                             X
  30.00%                                           X
  25.00%                                          X
  20.00%                                        X
  15.00%                                      X
  10.00%                                     X
   5.00%                                   X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}
