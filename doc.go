/*
Package Signals generates, stores, downloads and manipulates abstract signals, when imported it can then be used with specific real-world quantities.

###Definition of a 'signal'

>A varying value of some property, as it depends, uniquely, on some parameter.
>The controlling parameter is generally unbounded, and the property bounded.

see; https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.


###Fundamental Types

x :- the 'parameter' designed to be used as if it were unbounded (+ve and -ve), with unitX near the centre of its precision range.
y :- the 'property', a value between limits, +unitY and -unitY.

(the underlying types of x and y are kept hidden to enable simple generation of optimised packages with different ranges/precisions.)


###Interfaces

**Signal**

has one method, property, which returning a 'y' value from an 'x' value parameter

fundamentally procedural, calculated as needed, so that any 'x' value returns a 'y' value.

changes to parameters effect returned values from any other Signals composed from them.

saved/loaded, lossily, as PCM data. (PCM data can be encoded and saved in a Waveform Audio File Format (wav) file.)

saved/loaded from a go code binary (gob) file, (and signals can stream data, including gob files.) making for a basic interpreted signal language.

**LimitedSignal** 

a Signal with an additional method; MaxX(), that returns the 'x' value above which the Signal can be assumed to return zero, effectively the Signals end.

when required, an 'x' value of zero is regarded as a Signals start.

**PeriodicSignal** 

a Signal with an additional method; Period(), returning the 'x' length over which it repeats.

or when required any fundamental wavelength

or the sample spacing for one of the PCM Signal types.

**PeriodicLimitedSignal** :- both above, and is implemented by the PCM Signal types.

*/
package signals

/*
Implementation details.

x and y are not exported, separating their abstract nature from an importing packages concrete implementation and allowing flexibility in representation, if needed they can be made through provided exposed functions.
x and y are encoded as non-floating types, so resolution doesn't vary with value. By changing unitX the precision of a value can be directly effected, and the overall range can be altered, making for a basic ability to 'float' the range of the variable, just not automatically.

*/
