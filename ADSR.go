package signals

import	"encoding/gob"


func init() {
	gob.Register(ADSREnvelope{})
}

//  Attack Decay Sustain Release (ADSR) envelope.  see https://en.wikipedia.org/wiki/Synthesizer#Attack_Decay_Sustain_Release_.28ADSR.29_envelope
type ADSREnvelope struct {
	attackEnd    interval
	attackSlope  level
	decaySlope   level
	sustainStart interval
	sustain      level
	sustainEnd   interval
	releaseSlope level
	end          interval
}

func NewADSREnvelope(attack, decay, sustain interval, sustainlevel level, release interval) ADSREnvelope {
	// TODO release attack or decay of zero!
	return ADSREnvelope{attack, MaxLevel / level(attack), (MaxLevel - sustainlevel) / level(decay), attack + decay, sustainlevel, attack + decay + sustain, sustainlevel / level(release), attack + decay + sustain + release}
}

func (s ADSREnvelope) Level(t interval) level {
	if t > s.end {
		return 0
	} else if t > s.sustainEnd {
		return level(s.end-t) * s.releaseSlope
	} else if t > s.sustainStart {
		return s.sustain
	} else if t > s.attackEnd {
		return level(s.sustainStart-t)*s.decaySlope + s.sustain
	} else if t > 0 {
		return level(t) * s.attackSlope
	} else {
		return 0
	}
}
