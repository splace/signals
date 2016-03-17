package signals

import (
	"fmt"
	"strings"
)

func ExampleNoise() {
	s := NewNoise()
	for t := x(0); t < 40*unitX; t += unitX {
		fmt.Println("\t", s.call(t), strings.Repeat(" ", int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
	/* Output:
	     23.94%                                         X
	    -52.49%                 X
	      8.21%                                    X
	     -9.87%                               X
	    -74.46%          X
	    -68.54%            X
	    -31.13%                        X
	    -28.89%                         X
	     11.03%                                     X
	     43.01%                                                X
	    -71.97%           X
	    -35.88%                       X
	    -58.86%               X
	     47.80%                                                 X
	     21.68%                                         X
	    -34.58%                       X
	    -66.41%             X
	     10.38%                                     X
	      4.28%                                   X
	    -14.14%                              X
	    -17.82%                             X
	    -31.24%                        X
	     22.84%                                         X
	    -21.90%                           X
	     17.72%                                       X
	     23.27%                                         X
	     38.15%                                              X
	     65.67%                                                       X
	    -72.58%           X
	    -66.54%             X
	    -33.93%                       X
	      4.60%                                   X
	    -42.08%                     X
	    -36.43%                      X
	    -48.60%                  X
	    -10.65%                               X
	    -17.75%                             X
	     25.50%                                          X
	     23.76%                                         X
	    -87.69%      X
	*/
}

