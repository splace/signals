/*
Package signals generates and manipulates signals:- https://en.wikipedia.org/wiki/Signal_processing.

Overview

signals are entirely procedural, all values are calculated as needed, currently no cache.

this package supports only 1-Dimensionsal variation, and for simplicity the terminolology used represents analogue variation in time.

this package is intended to be general, and so a base package for import, and then used with specific real-world quantities.

Fundamental Types

Level :- can have a value from -MaxLevel to +MaxLevel

Interval :- 'usually' can be used as if infinite (+ve and -ve), with UnitTime somewhere near the center of its precision range.

Interfaces

Signal :- has method Level() which returns a Level value from an Interval value parameter.

Periodical :- a Signal with an additional method  Period(), that returns the signals assumed repeat period Interval.
	
*/
package signals

