package signals

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func PrintGraph(s Signal, start, end, step x) {
	for t := start; t < end; t += step {
		fmt.Println(s.property(t), strings.Repeat(" ", int(s.property(t)/(unitY/33))+33)+"X")
	}
}

func ExampleConstantZero() {
	PrintGraph(Constant{0}, 0, 3*unitX, unitX)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}

func ExampleConstantUnity() {
	PrintGraph(NewConstant(0), 0, 3*unitX, unitX)
	/* Output:
 100.00%                                                                  X
 100.00%                                                                  X
 100.00%                                                                  X
	*/
}

func ExampleSquare() {
	PrintGraph(Square{unitX}, 0, 2*unitX, unitX/10)
	/* Output:
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
-100.00% X
-100.00% X
-100.00% X
-100.00% X
-100.00% X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
-100.00% X
-100.00% X
-100.00% X
-100.00% X
-100.00% X
	*/
}

func ExamplePulse() {
	PrintGraph(Pulse{unitX}, -2*unitX, 3*unitX, unitX/4)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}
func ExampleRampUpDown() {
	PrintGraph(RampUp{unitX}, 0, 2*unitX, unitX/10)
	fmt.Println()
	PrintGraph(RampDown{unitX}, 0, 2*unitX, unitX/10)
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
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 
 100.00%                                                                  X
  90.00%                                                               X
  80.00%                                                            X
  70.00%                                                         X
  60.00%                                                     X
  50.00%                                                  X
  40.00%                                               X
  30.00%                                           X
  20.00%                                        X
  10.00%                                     X
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
func ExampleHeavyside() {
	PrintGraph(Heavyside{}, -3*unitX, 3*unitX, unitX/4)
	/* Output:
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
   0.00%                                  X
   0.00%                                  X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
	*/
}

func ExampleSine() {
	PrintGraph(Sine{unitX}, 0, 2*unitX, unitX/16)
	/* Output:
   0.00%                                  X
  38.27%                                              X
  70.71%                                                         X
  92.39%                                                                X
 100.00%                                                                  X
  92.39%                                                                X
  70.71%                                                         X
  38.27%                                              X
   0.00%                                  X
 -38.27%                      X
 -70.71%           X
 -92.39%    X
-100.00%  X
 -92.39%    X
 -70.71%           X
 -38.27%                      X
  -0.00%                                  X
  38.27%                                              X
  70.71%                                                         X
  92.39%                                                                X
 100.00%                                                                  X
  92.39%                                                                X
  70.71%                                                         X
  38.27%                                              X
   0.00%                                  X
 -38.27%                      X
 -70.71%           X
 -92.39%    X
-100.00%  X
 -92.39%    X
 -70.71%           X
 -38.27%                      X
	*/
}

func ExampleNewTone() {
	PrintGraph(NewTone(unitX, 0), 0, 2*unitX, unitX/16)
	/* Output:
   0.00%                                  X
  38.27%                                              X
  70.71%                                                         X
  92.39%                                                                X
 100.00%                                                                  X
  92.39%                                                                X
  70.71%                                                         X
  38.27%                                              X
   0.00%                                  X
 -38.27%                      X
 -70.71%           X
 -92.39%    X
-100.00%  X
 -92.39%    X
 -70.71%           X
 -38.27%                      X
  -0.00%                                  X
  38.27%                                              X
  70.71%                                                         X
  92.39%                                                                X
 100.00%                                                                  X
  92.39%                                                                X
  70.71%                                                         X
  38.27%                                              X
   0.00%                                  X
 -38.27%                      X
 -70.71%           X
 -92.39%    X
-100.00%  X
 -92.39%    X
 -70.71%           X
 -38.27%                      X
	*/
}

func ExampleSigmoid() {
	PrintGraph(Sigmoid{unitX}, -5*unitX, 5*unitX, unitX/2)
	/* Output:
   0.67%                                  X
   1.10%                                  X
   1.80%                                  X
   2.93%                                  X
   4.74%                                   X
   7.59%                                    X
  11.92%                                     X
  18.24%                                        X
  26.89%                                          X
  37.75%                                              X
  50.00%                                                  X
  62.25%                                                      X
  73.11%                                                          X
  81.76%                                                            X
  88.08%                                                               X
  92.41%                                                                X
  95.26%                                                                 X
  97.07%                                                                  X
  98.20%                                                                  X
  98.90%                                                                  X
	*/
}

func ExampleShifted() {
	PrintGraph(Shifted{NewADSREnvelope(unitX, unitX, unitX, unitY/2, unitX),unitX}, 0, 5*unitX, unitX/10)
	/* Output:
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
	*/
}

func ExampleReflected() {
	PrintGraph(Reflected{NewADSREnvelope(unitX, unitX, unitX, unitY/2, unitX)}, 0, 5*unitX, unitX/10)
	/* Output:
 100.00%                                                                   X
  90.00%                                                               X
  80.00%                                                            X
  70.00%                                                         X
  60.00%                                                     X
  50.00%                                                  X
  40.00%                                               X
  30.00%                                           X
  20.00%                                        X
  10.00%                                     X
   0.00%                                  X
   5.00%                                   X
  10.00%                                     X
  15.00%                                      X
  20.00%                                        X
  25.00%                                          X
  30.00%                                           X
  35.00%                                             X
  40.00%                                               X
  45.00%                                                X
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
  55.00%                                                    X
  60.00%                                                     X
  65.00%                                                       X
  70.00%                                                         X
  75.00%                                                          X
  80.00%                                                            X
  85.00%                                                              X
  90.00%                                                               X
  95.00%                                                                 X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
 100.00%                                                                   X
	*/
}

func ExamplePower() {
	Sine := Sine{unitX * 2}
	Power := Modulated{Sine, Sine}
	PrintGraph(Power, 0, unitX, unitX/40)
	/* Output:
   0.00%                                  X
   0.62%                                  X
   2.45%                                  X
   5.45%                                   X
   9.55%                                     X
  14.64%                                      X
  20.61%                                        X
  27.30%                                           X
  34.55%                                             X
  42.18%                                               X
  50.00%                                                  X
  57.82%                                                     X
  65.45%                                                       X
  72.70%                                                         X
  79.39%                                                            X
  85.36%                                                              X
  90.45%                                                               X
  94.55%                                                                 X
  97.55%                                                                  X
  99.38%                                                                  X
 100.00%                                                                  X
  99.38%                                                                  X
  97.55%                                                                  X
  94.55%                                                                 X
  90.45%                                                               X
  85.36%                                                              X
  79.39%                                                            X
  72.70%                                                         X
  65.45%                                                       X
  57.82%                                                     X
  50.00%                                                  X
  42.18%                                               X
  34.55%                                             X
  27.30%                                           X
  20.61%                                        X
  14.64%                                      X
   9.55%                                     X
   5.45%                                   X
   2.45%                                  X
   0.62%                                  X
	*/
}

func ExampleModulated() {
	PrintGraph(Modulated{Sine{unitX * 2}, Sine{unitX * 5}}, 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
   3.87%                                   X
  14.62%                                      X
  29.78%                                           X
  45.82%                                                 X
  58.78%                                                     X
  65.10%                                                       X
  62.34%                                                      X
  49.63%                                                  X
  27.96%                                           X
   0.00%                                  X
 -30.35%                        X
 -58.66%               X
 -80.74%        X
 -93.42%    X
 -95.11%   X
 -86.05%      X
 -68.31%            X
 -45.29%                    X
 -21.15%                            X
  -0.00%                                  X
  14.89%                                      X
  21.64%                                         X
  20.12%                                        X
  11.92%                                     X
   0.00%                                  X
 -11.92%                               X
 -20.12%                            X
 -21.64%                           X
 -14.89%                              X
   0.00%                                  X
  21.15%                                        X
  45.29%                                                X
  68.31%                                                        X
  86.05%                                                              X
  95.11%                                                                 X
  93.42%                                                                X
  80.74%                                                            X
  58.66%                                                     X
  30.35%                                            X
   0.00%                                  X
 -27.96%                         X
 -49.63%                  X
 -62.34%              X
 -65.10%             X
 -58.78%               X
 -45.82%                   X
 -29.78%                         X
 -14.62%                              X
  -3.87%                                 X
	*/
}

func ExampleStack() {
	PrintGraph(Stack{Sine{unitX * 2}, Sine{unitX * 5}}, 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
  21.72%                                         X
  41.82%                                               X
  58.86%                                                     X
  71.64%                                                         X
  79.39%                                                            X
  81.78%                                                            X
  78.98%                                                            X
  71.61%                                                         X
  60.69%                                                      X
  47.55%                                                 X
  33.66%                                             X
  20.51%                                        X
   9.45%                                     X
   1.56%                                  X
  -2.45%                                  X
  -2.31%                                  X
   1.77%                                  X
   9.14%                                     X
  18.78%                                        X
  29.39%                                           X
  39.54%                                               X
  47.80%                                                 X
  52.89%                                                   X
  53.82%                                                   X
  50.00%                                                  X
  41.29%                                               X
  28.02%                                           X
  10.98%                                     X
  -8.64%                                X
 -29.39%                         X
 -49.68%                  X
 -67.91%            X
 -82.67%       X
 -92.79%    X
 -97.55%  X
 -96.67%   X
 -90.35%     X
 -79.29%        X
 -64.57%             X
 -47.55%                   X
 -29.79%                         X
 -12.83%                              X
   1.93%                                  X
  13.33%                                      X
  20.61%                                        X
  23.47%                                         X
  22.04%                                         X
  16.95%                                       X
   9.18%                                     X
	*/
}

func ExampleTriggered() {
	s := NewTriggered(NewADSREnvelope(unitX, unitX, unitX, unitY/2, unitX), unitY/3*2, true, unitX/100, unitX*10)
	PrintGraph(s, 0, 5*unitX, unitX/10)
	fmt.Println(s.Found.Shift)
	s.Rising = false // forces a new search from here
	PrintGraph(s, 0, 5*unitX, unitX/10)
	fmt.Println(s.Found.Shift)
	/* Output:
  67.00%                                                        X
  77.00%                                                           X
  87.00%                                                              X
  97.00%                                                                  X
  96.50%                                                                 X
  91.50%                                                                X
  86.50%                                                              X
  81.50%                                                            X
  76.50%                                                           X
  71.50%                                                         X
  66.50%                                                       X
  61.50%                                                      X
  56.50%                                                    X
  51.50%                                                  X
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
  46.50%                                                 X
  41.50%                                               X
  36.50%                                              X
  31.50%                                            X
  26.50%                                          X
  21.50%                                         X
  16.50%                                       X
  11.50%                                     X
   6.50%                                    X
   1.50%                                  X
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
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
     0.67
  66.50%                                                       X
  61.50%                                                      X
  56.50%                                                    X
  51.50%                                                  X
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
  46.50%                                                 X
  41.50%                                               X
  36.50%                                              X
  31.50%                                            X
  26.50%                                          X
  21.50%                                         X
  16.50%                                       X
  11.50%                                     X
   6.50%                                    X
   1.50%                                  X
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
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
     1.67
	*/
}

func ExampleSegmented() {
	PrintGraph(NewSegmented(Sine{unitX * 10}, unitX), 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
   5.88%                                   X
  11.76%                                     X
  17.63%                                       X
  23.51%                                         X
  29.39%                                           X
  35.27%                                             X
  41.14%                                               X
  47.02%                                                 X
  52.90%                                                   X
  58.78%                                                     X
  62.41%                                                      X
  66.04%                                                       X
  69.68%                                                        X
  73.31%                                                          X
  76.94%                                                           X
  80.57%                                                            X
  84.21%                                                             X
  87.84%                                                              X
  91.47%                                                                X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  95.11%                                                                 X
  91.47%                                                                X
  87.84%                                                              X
  84.21%                                                             X
  80.57%                                                            X
  76.94%                                                           X
  73.31%                                                          X
  69.68%                                                        X
  66.04%                                                       X
  62.41%                                                      X
  58.78%                                                     X
  52.90%                                                   X
  47.02%                                                 X
  41.14%                                               X
  35.27%                                             X
  29.39%                                           X
  23.51%                                         X
  17.63%                                       X
  11.76%                                     X
   5.88%                                   X
	*/

}

func ExampleSegmented_makeSawtooth() {
	PrintGraph(NewSegmented(Square{unitX}, unitX/2), 0, 2*unitX, unitX/10)
	/* Output:
 100.00%                                                                  X
  60.00%                                                     X
  20.00%                                        X
 -20.00%                            X
 -60.00%               X
-100.00%  X
 -60.00%               X
 -20.00%                            X
  20.00%                                        X
  60.00%                                                     X
 100.00%                                                                  X
  60.00%                                                     X
  20.00%                                        X
 -20.00%                            X
 -60.00%               X
-100.00%  X
 -60.00%               X
 -20.00%                            X
  20.00%                                        X
  60.00%                                                     X
	*/
}

func ExampleRateModulated() {
	PrintGraph(RateModulated{Sine{unitX * 5}, Sine{unitX * 10}, unitX}, 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
  20.31%                                        X
  39.75%                                               X
  57.49%                                                    X
  72.78%                                                          X
  85.03%                                                              X
  93.79%                                                                X
  98.78%                                                                  X
  99.92%                                                                  X
  97.29%                                                                  X
  91.13%                                                                X
  81.82%                                                             X
  69.86%                                                         X
  55.80%                                                    X
  40.23%                                               X
  23.77%                                         X
   6.99%                                    X
  -9.57%                               X
 -25.46%                          X
 -40.26%                     X
 -53.69%                 X
 -65.52%             X
 -75.61%          X
 -83.90%       X
 -90.38%     X
 -95.11%   X
 -98.18%  X
 -99.74%  X
 -99.92%  X
 -98.89%  X
 -96.83%   X
 -93.88%    X
 -90.22%     X
 -85.99%      X
 -81.32%        X
 -76.32%         X
 -71.11%           X
 -65.76%             X
 -60.34%               X
 -54.91%                X
 -49.51%                  X
 -44.18%                    X
 -38.93%                      X
 -33.78%                       X
 -28.73%                         X
 -23.77%                           X
 -18.90%                            X
 -14.10%                              X
  -9.37%                               X
  -4.67%                                 X
	*/
}

func ExampleLooped() {
	PrintGraph(Looped{Sine{unitX * 5}, unitX * 25 / 10}, 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
  12.53%                                      X
  24.87%                                          X
  36.81%                                              X
  48.18%                                                 X
  58.78%                                                     X
  68.45%                                                        X
  77.05%                                                           X
  84.43%                                                             X
  90.48%                                                               X
  95.11%                                                                 X
  98.23%                                                                  X
  99.80%                                                                  X
  99.80%                                                                  X
  98.23%                                                                  X
  95.11%                                                                 X
  90.48%                                                               X
  84.43%                                                             X
  77.05%                                                           X
  68.45%                                                        X
  58.78%                                                     X
  48.18%                                                 X
  36.81%                                              X
  24.87%                                          X
  12.53%                                      X
   0.00%                                  X
  12.53%                                      X
  24.87%                                          X
  36.81%                                              X
  48.18%                                                 X
  58.78%                                                     X
  68.45%                                                        X
  77.05%                                                           X
  84.43%                                                             X
  90.48%                                                               X
  95.11%                                                                 X
  98.23%                                                                  X
  99.80%                                                                  X
  99.80%                                                                  X
  98.23%                                                                  X
  95.11%                                                                 X
  90.48%                                                               X
  84.43%                                                             X
  77.05%                                                           X
  68.45%                                                        X
  58.78%                                                     X
  48.18%                                                 X
  36.81%                                              X
  24.87%                                          X
  12.53%                                      X
	*/
}

func ExampleRepeated() {
	PrintGraph(Repeated{Sine{unitX * 2}, 1.5}, 0, 5*unitX, unitX/10)
	/* Output:
   0.00%                                  X
  30.90%                                            X
  58.78%                                                     X
  80.90%                                                            X
  95.11%                                                                 X
 100.00%                                                                  X
  95.11%                                                                 X
  80.90%                                                            X
  58.78%                                                     X
  30.90%                                            X
   0.00%                                  X
 -30.90%                        X
 -58.78%               X
 -80.90%        X
 -95.11%   X
-100.00%  X
 -95.11%   X
 -80.90%        X
 -58.78%               X
 -30.90%                        X
   0.00%                                  X
  30.90%                                            X
  58.78%                                                     X
  80.90%                                                            X
  95.11%                                                                 X
 100.00%                                                                  X
  95.11%                                                                 X
  80.90%                                                            X
  58.78%                                                     X
  30.90%                                            X
   0.00%                                  X
  30.90%                                            X
  58.78%                                                     X
  80.90%                                                            X
  95.11%                                                                 X
 100.00%                                                                  X
  95.11%                                                                 X
  80.90%                                                            X
  58.78%                                                     X
  30.90%                                            X
   0.00%                                  X
 -30.90%                        X
 -58.78%               X
 -80.90%        X
 -95.11%   X
-100.00%  X
 -95.11%   X
 -80.90%        X
 -58.78%               X
 -30.90%                        X
	*/
}

func BenchmarkSignalsSine(b *testing.B) {
	b.StopTimer()
	s := Sine{unitX}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Encode(ioutil.Discard, s, unitX, 44100, 1)
	}

}

func BenchmarkSignalsSineSegmented(b *testing.B) {
	b.StopTimer()
	s := NewSegmented(Sine{unitX}, unitX/512)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Encode(ioutil.Discard, s, unitX, 44100, 1)
	}

}


