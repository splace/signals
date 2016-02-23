/*
Package signals generates and manipulates abstract signals:- https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.

Overview

Functions are procedural, y's are calculated as needed, meaning changes in parameters, or arrangment, are immediately effective.

PCMFunctions are stored, at a particular interval and precision, and can be used to cache an expensive precedural Function.

intended to be abstract, and a base package for import, then used with specific real-world quantities.

Functions can be encode/decoded as go code binary (gob), making for a basic interpreted signal language.

or they can be, lossily, stored in wav files (functions saved as wav are loaded back as PCMFunctions)

Fundamental Types

x :- 'usually' can be used as if infinite (+ve and -ve), with UnitX somewhere near the center of its precision range.

y :- can have a value from -Maxy to +Maxy

Interfaces

Function :- has method Call() which returns a y value from an x value parameter.

LimitedFunction :- has a MaxX() method that returns the x value after which the function can be assumed to return zero. ie ends.

PeriodicFunction :- a Function with an additional method Period(), the repeat, (or sampling for PCM), delta x.

PCMFunction :- a PeriodicLimitedFunction with additional method Encode().
(NB Period() for PCMFunctions represents the sample spacing, since repeating stored functions are assumed not required, use a Looped modifier.)

*/
package signals
