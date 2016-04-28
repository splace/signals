/*
Package signals generates and manipulates abstract signals, when imported can then be used with specific real-world quantities.


Definition of signal

A signal is the record of some property, as it varies with some parameter.

The controlling parameter is generally unboundeed, and the property bounded.

also see; https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.


Fundamental Types

x :- the 'parameter' designed to be used as if unbounded (+ve and -ve), with unitX near the center of its precision range.

y :- the 'property', can have a value between +unitY and -unitY.


Interfaces

Function :- has method property(x), which returns a 'y' value from an 'x' value parameter.
Function's are generally procedural, calculated as needed, meaning changes in parameters, or arrangment, effect returned values of existing types.
Function's can be encode/decoded as go code binary (gob), making for a basic interpreted signal language, or they can be stored, lossily, in wav files (Function's saved as wav are loaded back as PCMFunctions)

LimitedFunction :- has a MaxX() method that returns the 'x' value after which the function can be assumed to return zero, effectively has an end.
an 'x' value of zero is normally requarded as its start.

PeriodicFunction :- a Function with an additional method Period(), the repeat dx, r the reciprocal of any fundamental frequency, or the sample spacing for PCMFunction's.

PeriodicLimitedFunction :- both above.

PCMFunction :- a PeriodicLimitedFunction with additional method Encode().
PCMFunction's are stored, at a particular interval and precision, and can be used to cache an expensive precedural Function.


*/
package signals

/*
Implementation details.

x and y are not exported, separating their abstract nature from an importing packages concrete implementation and allowing flexability in representation, if needed they can be made through provided exposed functions.
x and y are encoded as fixed precision, so resolution doesn't vary with value.
*/
