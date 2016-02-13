/*
Package signals generates and manipulates signals:- https://en.wikipedia.org/wiki/Signal_processing.

signals are entirely procedural.

signals "any quantity exhibiting variation in time or variation in space".

currently this package supports only 1-Dimensionsal variation, and for simplicity terminolology used represents analogue variation in time.

this package is intended to be general, and so a base package for import, and used then with specific real-world quantities.

Interfaces

	Signal :- method Level() returns a Level value from an Interval value parameter.
	Tone :- a Signal with a method  Period() that returns its repeat period Interval.
	
Types

	Level :- a value from -MaxLevel to +MaxLevel
	Interval :- a value with UnitTime somewhere near the center of its range. 
*/
package signals
/*
doc test
*/
