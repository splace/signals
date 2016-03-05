/*
Package signals generates and manipulates abstract signals, often imported and then used with specific real-world quantities.

Signals variously defined as:

"A detectable physical quantity or impulse (as a voltage, current, or magnetic field strength) by which messages or information can be transmitted." or

"A signal is a function of independent variables that carry some information."

"A signal is a source of information generaly a physical quantity which varies with respect to time, space, temperature like any independent variable"

"A signal is a physical quantity that varies with time,space or any other independent variable.by which information can be conveyed"

see; https://en.wikibooks.org/wiki/Signals_and_Systems/Definition_of_Signals_and_Systems.


Overview

Interfaces

Function :- has method call() which returns a 'y' value from an 'x' value parameter.
Function's are generally procedural, calculated as needed, meaning changes in parameters, or arrangment, effect returned values of existing types.
Function's can be encode/decoded as go code binary (gob), making for a basic interpreted signal language, or they can be stored, lossily, in wav files (Function's saved as wav are loaded back as PCMFunctions)

LimitedFunction :- has a MaxX() method that returns the 'x' value after which the function can be assumed to return zero, effectively has an end.

PeriodicFunction :- a Function with an additional method Period(), the repeat dx, reciprocal of any fundamental frequency, or sample spacing for PCMFunction's.

PCMFunction :- a PeriodicLimitedFunction with additional method Encode().
PCMFunction's are stored, at a particular interval and precision, and can be used to cache an expensive precedural Function.


Fundamental Types

x :- 'usually' can be used as if infinite (+ve and -ve), with UnitX somewhere near the center of its precision range.

y :- can have a value from -Maxy to +Maxy

Note: x and y are package local to allow changes to how they are represented internally to be flexible, some access is provided through functions.
*/
package signals
