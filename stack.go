package signals

import (
	"encoding/gob"
	"io"
)

func init() {
	gob.Register(Stack{})
	gob.Register(Sum{})
}

// Stack is a Signal generated by adding together Signal(s).
// source Signals are scaled down by Stacks  count, making it impossible to overrun maxlevel.
// also a Periodical, taking its period, if any, from its first member.
// like OR logic, all sources have to be zero (at a particular momemt) for Stack to be zero.
type Stack []Signal

func (c Stack) Level(t interval) (total level) {
	for _, s := range c {
		total += s.Level(t)/level(len(c))
	}
	return
}

func (c Stack) Period() (period interval) {
	if len(c) > 0 {
		if s, ok := c[0].(Periodical); ok {
			return s.Period()
		}
	}
	return
}

func (c Stack) Save(p io.Writer) error {
	return gob.NewEncoder(p).Encode(&c)
}

func (c *Stack) Load(p io.Reader) error {
	return gob.NewDecoder(p).Decode(c)
}
// helper: needed becasue can't use type literal with array source. 
func NewStack(c ...Signal) Stack{
	return Stack(c)
}


// Sum is a Stack that doesn't scale its contents, so can overflow.
type Sum Stack

func (c Sum) Level(t interval) (total level) {
	for _, s := range c {
		total += s.Level(t)
	}
	return
}

// helper: needed becasue can't use type literal with array source. 
func NewSum(c ...Signal) Sum{
	return Sum(c)
}


