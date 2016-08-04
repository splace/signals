/*
Package signals generates, stores, downloads and manipulates abstract signals, when imported it can then be used with specific real-world quantities.


Definition of 'signal'

A varying value of some property, as it depends, uniquely, on some parameter.

The controlling parameter is generally unbounded, and the property bounded.

also see; https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.


Fundamental Types

x :- 'parameter' designed to be used as if unbounded (+ve and -ve), with unitX near the centre of its precision range.

y :- 'property', can have a value between limits, +unitY and -unitY.


Interfaces

Signal :- has method property(x), which returns a 'y' value from an 'x' value parameter.
Signal's are generally procedural, calculated as needed, meaning changes in parameters, or arrangement, effect returned values of existing types.
Signal's can be saved/loaded from a go code binary (gob) file, (and signals can stream data, including gob files.) making for a basic interpreted signal language, or they can be stored, lossily, as PCM data. (PCM data can be encoded and saved in a Waveform Audio File Format (wav) file.)

LimitedSignal :- has a MaxX() method that returns the 'x' value above which the Signal can be assumed to return zero, effectively has an end.
an 'x' value of zero is normally regarded as its start.

PeriodicSignal :- a Signal with an additional method Period(), returning the repeat dx, or any fundamental wavelength, or the sample spacing for one of te PCM Signal types.

PeriodicLimitedSignal :- both above, and is implemented by the PCM Signal types.

*/
package signals

/*
Implementation details.

x and y are not exported, separating their abstract nature from an importing packages concrete implementation and allowing flexibility in representation, if needed they can be made through provided exposed functions.
x and y are encoded as non-floating types, so resolution doesn't vary with value, but by changing unitX the precision of a value can be changed, and the overall range can be altered, making for a basic ability to 'float' the range of the variable, but only when directed, not automatically.

*/
