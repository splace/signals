/*
Package signals generates and manipulates abstract signals, when imported can then be used with specific real-world quantities.


Definition of signal

A signal is the value of some property, as it depends, uniquely, on some parameter.

The controlling parameter is generally unbounded, and the property bounded.

also see; https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.


Fundamental Types

x :- the 'parameter' designed to be used as if unbounded (+ve and -ve), with unitX near the centre of its precision range.

y :- the 'property', can have a value between +unitY and -unitY.


Interfaces

Signal :- has method property(x), which returns a 'y' value from an 'x' value parameter.
Signal's are generally procedural, calculated as needed, meaning changes in parameters, or arrangement, effect returned values of existing types.
Signal's can be saved/loaded from go code binary (gob), making for a basic interpreted signal language, or they can be stored, lossily, in wav files (Signal's saved as wav are loaded back as PCMSignals)

LimitedSignal :- has a MaxX() method that returns the 'x' value after which the Signal can be assumed to return zero, effectively has an end.
an 'x' value of zero is normally regarded as its start.

PeriodicSignal :- a Signal with an additional method Period(), returning the repeat dx, or the reciprocal of any fundamental frequency, or the sample spacing for PCMSignal's.

PeriodicLimitedSignal :- both above.

PCMSignal :- a PeriodicLimitedSignal with additional method Encode().
PCMSignal's are stored 'recordings' rather than procedurally generated, at a particular interval and precision. They can be used to cache a procedural Signal.


*/
package signals

/*
Implementation details.

x and y are not exported, separating their abstract nature from an importing packages concrete implementation and allowing flexibility in representation, if needed they can be made through provided exposed functions.
x and y are encoded as fixed precision, so resolution doesn't vary with value.
*/
