package signals

import (
	"encoding/gob"
	"math/rand"
	"sync"
)

func init() {
	gob.Register(Noise{})
}

// Noise is a deterministic random Signal, white noise.
// it always produces the same y value for the same x value, (for the same Noise) but random otherwise.
// determinism allows caching even for this type
type Noise struct {
	rand.Rand
	sync.Mutex
}

func NewNoise() Noise {
	return Noise{Rand:*rand.New(rand.NewSource(rand.Int63()))} // give each noise, very probably, a different random source
}

func (s Noise) property(p x) (v y) {
	s.Lock()
	rand.Seed(int64(p)) // set the default generators seed to the same, for the same x
	s.Seed(int64(rand.Int63())) // a Noise sets its generator's seed to a random number from the default generator, which will be the same for the same x, and so the same random numbers will be generated from it, but will be different for different Noises.
	v = y(s.Int63())
	v -= y(s.Int63())  // makes for normal dist (linear delta) and full range
	s.Unlock()
	return
}
