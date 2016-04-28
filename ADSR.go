package signals

import "encoding/gob"

func init() {
	gob.Register(ADSREnvelope{})
}

//  Attack Decay Sustain Release (ADSR) envelope.  see https://en.wikipedia.org/wiki/Synthesizer#Attack_Decay_Sustain_Release_.28ADSR.29_envelope
type ADSREnvelope struct {
	attackEnd    x
	attackSlope  y
	decaySlope   y
	sustainStart x
	sustain      y
	sustainEnd   x
	releaseSlope y
	end          x
}

func NewADSREnvelope(attack, decay, sustain x, sustainy y, release x) ADSREnvelope {
	// TODO release attack or decay of zero!
	return ADSREnvelope{attack, unitY / y(attack), (unitY - sustainy) / y(decay), attack + decay, sustainy, attack + decay + sustain, sustainy / y(release), attack + decay + sustain + release}
}

func (s ADSREnvelope) property(t x) y {
	if t > s.end {
		return 0
	} else if t > s.sustainEnd {
		return y(s.end-t) * s.releaseSlope
	} else if t > s.sustainStart {
		return s.sustain
	} else if t > s.attackEnd {
		return y(s.sustainStart-t)*s.decaySlope + s.sustain
	} else if t > 0 {
		return y(t) * s.attackSlope
	} else {
		return 0
	}
}

func (s ADSREnvelope) MaxX() x {
	return s.end
}
