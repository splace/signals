package signals

import (
	"math/rand"
	"encoding/gob"
)

func init() {
	gob.Register(Noise{})
}

// Noise is a deterministic random level Signal, white noise.
// it has the same value at the same time, but random otherwise.
type Noise struct {
	generator rand.Rand
}

func NewNoise() Noise {
	return Noise{*rand.New(rand.NewSource(rand.Int63()))}  // give each noise, very probably, a different generator source
}


func (s Noise) Level(t interval) (l level) {
	rand.Seed(int64(t))  // default generator set to the same seed for the same time
	s.generator.Seed(int64(rand.Int63()))   // Noise sets its generator's seed to a random number from default generator, which is the same at a given t, so the same random numbers generated from it, for the same t, but different for different Noises.
	l+=level(s.generator.Int63())
	l-=level(s.generator.Int63())
	return
}






