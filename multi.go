package signals

import (
	"encoding/gob"
	"io"
)
func init() {
	gob.Register(Multiplex{})
}

// Multiplex is a Signal generated by multiplying together Signal(s).
// multiplication scales so that, MaxLevel*MaxLevel=MaxLevel (so MaxLevel is unity).
// like logic AND; all its signals (at a particular momemt) need to be MaxLevel to produce a Multi of MaxLevel, whereas, ANY signal at zero will generate a Multi of zero.
// Multiplex is also a Periodical, taking its period, if any, from its first member.
type Multiplex []Signal

func (c Multiplex) Level(t interval) (total level) {
	total = MaxLevel
	for _, s := range c {
		l := s.Level(t)
		switch l {
		case 0:
			total = 0
			break
		case MaxLevel:
			continue
		default:
			//total = (total / HalfLevel) * (l / HalfLevel)*2
			total = (total >> HalfLevelBits) * (l >> HalfLevelBits) * 2
		}
	}
	return
}

func (c Multiplex) Period() (period interval) {
	if len(c) > 0 {
		if s, ok := c[0].(Periodical); ok {
			return s.Period()
		}
	}
	return
}

func (c Multiplex) Save(p io.Writer) error {
	return gob.NewEncoder(p).Encode(&c)
}

func (c *Multiplex) Load(p io.Reader) error {
	return gob.NewDecoder(p).Decode(c)
}

// helper: needed becasue can't use type literal with array source. 
func NewMultiplex(c ...Signal) Multiplex{
	return Multiplex(c)
}

