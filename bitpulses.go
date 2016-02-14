package signals

import (
	"math/big"
)

// pulse train specified by the bits of a big int
type PulsePattern struct {
	BitPattern big.Int
	PulseWidth interval
}

func (s PulsePattern) Level(t interval) level {
	//if bp := s.BitPattern.BitLen() - int(t/s.PulseWidth); bp < 0 { // higher number bits come first
	if bp := int(t / s.PulseWidth); bp < 0 {
		return 0
	} else {
		if s.BitPattern.Bit(bp) == 1 {
			return MaxLevel
		} else {
			return 0
		}
	}
}
