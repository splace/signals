package signals

import (
	"fmt"
	"strings"
	"testing"
	"io/ioutil"
	)
	


func ExampleSquare() {
	s := Square{unitX}
	for t := X(0); t < X(2); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
	s := Pulse{unitX}
	for t := X(-2); t < X(3); t += unitX / 4 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
 */}
func ExampleRampUpDown() {
	s := RampUp{unitX}
	for t := X(0); t < X(2); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
	s2 := RampDown{unitX}
	for t := X(0); t < X(2); t += unitX / 10 {
		fmt.Println(s2.call(t),strings.Repeat(" ",int(s2.call(t)/(maxY/33))+33)+"X")
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
*/}
func ExampleHeavyside() {
	s := Heavyside{}
	for t := X(-3); t < X(3); t += unitX / 4 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
*/}

func ExampleSine() {
	s := Sine{unitX}
	for t := X(0); t < X(2); t += unitX / 16 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleNewTone() {
	s := NewTone(unitX, 0)
	for t := X(0); t < X(2); t += unitX / 16 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleSigmoid() {
	s := Sigmoid{unitX}
	for t := X(-5); t < X(5); t += unitX / 2 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleReflected() {
	s := Reflected{NewADSREnvelope(unitX, unitX, unitX, maxY/2, unitX)}
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleMultiplex() {
	s := Multiplex{Sine{unitX * 2}, Sine{unitX * 5}}
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleStack() {
	s := Stack{Sine{unitX * 2}, Sine{unitX * 5}}
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

func ExampleTriggered() {
	s := NewTriggered(NewADSREnvelope(unitX, unitX, unitX, maxY/2, unitX), maxY / 3 * 2, true, unitX / 100, unitX * 10)
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
	fmt.Println(s.Found.Shift)
	//s.Trigger = Maxy / 3
	s.Rising = false
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
	fmt.Println(s.Found.Shift)
	fmt.Println()
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
*/}


func ExampleSegmented() {
	s := NewSegmented(Sine{unitX * 10},unitX)
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
	s := NewSegmented(Square{unitX},unitX/2)
	for t := X(0); t < X(2); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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

func ExampleModulated() {
	s := Modulated{Sine{unitX * 5}, Sine{unitX * 10},unitX}
	for t := X(0); t < X(5); t += unitX / 10 {
		fmt.Println(s.call(t),strings.Repeat(" ",int(s.call(t)/(maxY/33))+33)+"X")
	}
	fmt.Println()
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
  */}

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
	s := NewSegmented(Sine{unitX},unitX/2)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Encode(ioutil.Discard, s, unitX, 44100, 1)
	}	

}


/*  Hal3 Sat Mar 5 22:44:05 GMT 2016 go version go1.5.1 linux/amd64
=== RUN   TestNoiseSave
--- PASS: TestNoiseSave (0.82s)
=== RUN   TestSaveLoad
--- PASS: TestSaveLoad (0.00s)
=== RUN   TestSaveWav
--- PASS: TestSaveWav (0.00s)
=== RUN   TestLoad
--- PASS: TestLoad (0.01s)
=== RUN   TestLoadChannels
--- PASS: TestLoadChannels (0.08s)
=== RUN   TestStackPCMs
--- PASS: TestStackPCMs (0.08s)
=== RUN   TestMultiplexTones
--- PASS: TestMultiplexTones (0.02s)
=== RUN   TestSaveLoadSave
--- PASS: TestSaveLoadSave (0.04s)
=== RUN   TestPiping
--- PASS: TestPiping (0.00s)
=== RUN   TestImagingSine
--- PASS: TestImagingSine (0.27s)
=== RUN   TestImaging
--- PASS: TestImaging (0.31s)
=== RUN   TestComposable
--- PASS: TestComposable (1.46s)
=== RUN   TestStackimage
--- PASS: TestStackimage (0.91s)
=== RUN   TestMultiplexImage
--- PASS: TestMultiplexImage (0.90s)
=== RUN   ExampleADSREnvelope
--- PASS: ExampleADSREnvelope (0.00s)
=== RUN   ExamplePulsePattern
--- PASS: ExamplePulsePattern (0.00s)
=== RUN   ExampleNoise
--- PASS: ExampleNoise (0.00s)
=== RUN   ExampleSquare
--- PASS: ExampleSquare (0.00s)
=== RUN   ExamplePulse
--- PASS: ExamplePulse (0.00s)
=== RUN   ExampleRampUpDown
--- PASS: ExampleRampUpDown (0.00s)
=== RUN   ExampleHeavyside
--- PASS: ExampleHeavyside (0.00s)
=== RUN   ExampleSine
--- PASS: ExampleSine (0.00s)
=== RUN   ExampleNewTone
--- PASS: ExampleNewTone (0.00s)
=== RUN   ExampleSigmoid
--- PASS: ExampleSigmoid (0.00s)
=== RUN   ExampleReflected
--- PASS: ExampleReflected (0.00s)
=== RUN   ExampleMultiplex
--- PASS: ExampleMultiplex (0.00s)
=== RUN   ExampleStack
--- PASS: ExampleStack (0.00s)
=== RUN   ExampleTriggered
--- PASS: ExampleTriggered (0.00s)
=== RUN   ExampleSegmented
--- PASS: ExampleSegmented (0.00s)
=== RUN   ExampleSegmented_makeSawtooth
--- PASS: ExampleSegmented_makeSawtooth (0.00s)
=== RUN   ExampleModulated
--- PASS: ExampleModulated (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/signals	4.928s
Sat Mar 5 22:44:11 GMT 2016 */
/*  Hal3 Sat Mar 5 22:44:56 GMT 2016 go version go1.5.1 linux/amd64
PASS
BenchmarkSignalsSine-2         	     500	   3148458 ns/op
BenchmarkSignalsSineSegmented-2	     500	   3874923 ns/op
ok  	_/home/simon/Dropbox/github/working/signals	4.236s
Sat Mar 5 22:45:02 GMT 2016 */

