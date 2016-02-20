/*
Package signals generates and manipulates signals:- https://en.wikipedia.org/wiki/Signal_processing.

Overview

Signals are procedural, levels are calculated as needed, meaning changes in parameters, or arrangment, are immediately effective.

PCMSignals are stored, at a particular precision, and can be used to cache an expensive precedural Signal. 

only 1-Dimensionsal variation, and for simplicity the terminolology used represents analogue variation in time.

intended to be general, and a base package for import, then used with specific real-world quantities.

Signals can be encode/decoded as go code binary (gob), (probably best not used for PCMSignals, where saving as wav files is available.)

Fundamental Types

level :- can have a value from -MaxLevel to +MaxLevel

interval :- 'usually' can be used as if infinite (+ve and -ve), with UnitTime somewhere near the center of its precision range.

Interfaces

Signal :- has method Level() which returns a Level value from an Interval value parameter.

PCMSignal :- a Signal with with additional methods; SamplePeriod() returns the interval of each sample and Duration() returns the total saved interval of the Signal. 

Periodical :- a Signal with an additional method Period(), that returns the signals assumed repeat period Interval.

Notes

PCMSignals are an evenly spaced array of levels, with different types that store at a particular precisions. 

*/
package signals
