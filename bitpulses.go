package signals

import "encoding/gob"

import (
	"math/big"
)

func init() {
	gob.Register(PulsePattern{})
}

// pulse train specified by the bits of a big int.
// littleendian, ignores high zero bits for MaxX().
type PulsePattern struct {
	BitPattern big.Int
	PulseWidth x
}

func (s PulsePattern) property(t x) y {
	//if bp := s.BitPattern.BitLen() - int(t/s.PulseWidth); bp < 0 { // higher number bits come first
	if bp := int(t / s.PulseWidth); bp < 0 {
		return 0
	} else {
		if s.BitPattern.Bit(bp) == 1 {
			return unitY
		} else {
			return 0
		}
	}
}

func (s PulsePattern) Period() x {
	return s.PulseWidth
}

func (s PulsePattern) MaxX() x {
	return s.PulseWidth * x(s.BitPattern.BitLen())
}
