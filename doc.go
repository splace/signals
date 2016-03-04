/*
Package signals generates and manipulates abstract signals:- https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.

Overview

intended to be abstract, and a base package for import, then used with specific real-world quantities.

Interfaces

Function :- has method Call() which returns a 'y' value from an 'x' value parameter.
Function's are generally procedural, calculated as needed, meaning changes in parameters, or arrangment, effect returned values of existing types.
Function's can be encode/decoded as go code binary (gob), making for a basic interpreted signal language, or they can be, lossily, stored in wav files (Function's saved as wav are loaded back as PCMFunctions)

LimitedFunction :- has a MaxX() method that returns the 'x' value after which the function can be assumed to return zero, effectively has an end.

PeriodicFunction :- a Function with an additional method Period(), reciprocal of any fundamental frequency, (or sample spacing for PCMFunction's), delta 'x'.

PCMFunction :- a PeriodicLimitedFunction with additional method Encode().
PCMFunction's are stored, at a particular interval and precision, and can be used to cache an expensive precedural Function.


Fundamental Types

x :- 'usually' can be used as if infinite (+ve and -ve), with UnitX somewhere near the center of its precision range.

y :- can have a value from -Maxy to +Maxy

Note: x and y are package local to allow changes to how they are represented internally to be flexible, some access is provided through functions.
*/
package signals
