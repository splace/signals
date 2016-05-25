package signals

import (
	"encoding/gob"
	"math/rand"
)

func init() {
	gob.Register(Noise{})
}

// Noise is a deterministic random Signal, white noise.
// it always produces the same y value for the same x value, (for the same Noise) but random otherwise.
// determinism allows caching even for this type
type Noise struct {
	generator rand.Rand
}

func NewNoise() Noise {
	return Noise{*rand.New(rand.NewSource(rand.Int63()))} // give each noise, very probably, a different generator source
}

func (s Noise) property(t x) (l y) {
	rand.Seed(int64(t))                   // default generator set to the same seed for the same x
	s.generator.Seed(int64(rand.Int63())) // Noise sets its generator's seed to a random number from default generator, which is the same at a given x, so the same random numbers generated from it, for the same x, but different for different Noises.
	l += y(s.generator.Int63())
	l -= y(s.generator.Int63())
	return
}
