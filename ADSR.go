package signals

//  Attack Decay Sustain Release (ADSR) envelope.  see https://en.wikipedia.org/wiki/Synthesizer#Attack_Decay_Sustain_Release_.28ADSR.29_envelope
type ADSREnvelope struct {
	attackEnd    Interval
	attackSlope  Level
	decaySlope   Level
	sustainStart Interval
	sustain      Level
	sustainEnd   Interval
	releaseSlope Level
	end          Interval
}

func NewADSREnvelope(attack, decay, sustain Interval, level Level, release Interval) ADSREnvelope {
	// TODO release attack or decay of zero!
	return ADSREnvelope{attack, MaxLevel / Level(attack), (MaxLevel - level) / Level(decay), attack + decay, level, attack + decay + sustain, level / Level(release), attack + decay + sustain + release}
}

func (s ADSREnvelope) Level(t Interval) Level {
	if t > s.end {
		return 0
	} else if t > s.sustainEnd {
		return Level(s.end-t) * s.releaseSlope
	} else if t > s.sustainStart {
		return s.sustain
	} else if t > s.attackEnd {
		return Level(s.sustainStart-t)*s.decaySlope + s.sustain
	} else if t > 0 {
		return Level(t) * s.attackSlope
	} else {
		return 0
	}
}
