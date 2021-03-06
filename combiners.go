package signals

import (
	"encoding/gob"
)

func init() {
	gob.Register(Modulated{})
	gob.Register(Composite{})
	gob.Register(Stacked{})
}

// Modulated is a PeriodicLimitedSignal, generated by multiplying together Signal(s).(Signal's can be PeriodicLimitedSignal's, so this can be hierarchical.)
// Multiplication scales so that, unitY*unitY=unitY.
// Modulated's MaxX() comes from the smallest contstituent MaxX(), (0 if none of the contained Signals are LimitedSignals.)
// Modulated's Period() comes from its first member.
// As with 'AND' logic, all sources have to be unitY (at a particular x) for Modulated to be unitY, whereas, ANY Signal at zero will generate a Modulated of zero.
type Modulated []Signal

func (c Modulated) property(p x) (total y) {
	total = unitY
	for _, s := range c {
		l := s.property(p)
		switch l {
		case 0:
			total = 0
			break
		case unitY:
			continue
		default:
			//total = (total / Halfy) * (l / Halfy)*2
			total = (total >> halfyBits) * (l >> halfyBits) * 2
		}
	}
	return
}

func (c Modulated) Period() (period x) {
	// TODO could helpfully be the shortest period and any constituent.
	if len(c) > 0 {
		if s, ok := c[0].(PeriodicSignal); ok {
			return s.Period()
		}
	}
	return
}

// the smallest Max X of the constituents.
func (c Modulated) MaxX() (min x) {
	for _, s := range c {
		if sls, ok := s.(LimitedSignal); ok {
			if newmin := sls.MaxX(); newmin > 0 && (min == 0 || newmin < min) {
				min = newmin
			}
		}
	}
	return
}

// helper to enable generation from another slice.
// will in general need to use a slice interface promoter function.
func NewModulated(c ...Signal) Modulated {
	return Modulated(c)
}

// Composite is a PeriodicLimitedSignal, generated by adding together Signal(s). (PeriodicLimitedSignal's are Signal's so this can be hierarchical.)
// Composite's MaxX() comes from the largest contstituent MaxX(), (0 if none of the contained Signals are LimitedSignals.)
// Composite's Period() comes from its first member.
// As with 'OR' logic, all sources have to be zero (at a particular x) for Composite to be zero.
type Composite []Signal

func (c Composite) property(p x) (total y) {
	for _, s := range c {
		total += s.property(p)
	}
	return
}

func (c Composite) Period() (period x) {
	// TODO could be longest period muliple common to all constituents.
	if len(c) > 0 {
		if s, ok := c[0].(PeriodicSignal); ok {
			return s.Period()
		}
	}
	return
}

// the largest Max X of the constituents.
func (c Composite) MaxX() (max x) {
	for _, s := range c {
		if sls, ok := s.(LimitedSignal); ok {
			if newmax := sls.MaxX(); newmax > max {
				max = newmax
			}
		}
	}
	return
}

// helper to enable generation from another slice.
// will in general need to use a slice interface promoter function.
func NewComposite(c ...Signal) Composite {
	return Composite(c)
}

// Same as Composite except that Stacked scales down by the number of signals, making it impossible to exceed unitY.
type Stacked []Signal

func (c Stacked) property(p x) (total y) {
	for _, s := range c {
		total += s.property(p) / y(len(c))
	}
	return
}

func (c Stacked) Period() (period x) {
	// TODO could be longest period muliple common to all constituents.
	if len(c) > 0 {
		if s, ok := c[0].(PeriodicSignal); ok {
			return s.Period()
		}
	}
	return
}

// the largest Max X of the constituents.
func (c Stacked) MaxX() (max x) {
	for _, s := range c {
		if sls, ok := s.(LimitedSignal); ok {
			if newmax := sls.MaxX(); newmax > max {
				max = newmax
			}
		}
	}
	return
}

// helper to enable generation from another slice.
// will in general need to use a slice interface promoter function.
func NewStacked(c ...Signal) Stacked {
	return Stacked(c)
}


// Sequenced is a LimitedSignal, generated by appending together LimitedSignal(s).
type Sequenced []LimitedSignal

func (c Sequenced) property(p x) y {
	for _, s := range c {
		if l:=s.MaxX();p-l<0{
			return s.property(p)
		}else{
			p -= l
		}
	}
	return 0
}

// sum of all MaxX's in slice.
func (c Sequenced) MaxX() (max x) {
	for _, s := range c {
		max+=s.MaxX()
	}
	return
}

func NewSequence(c ...LimitedSignal) Sequenced {
	return Sequenced(c)
}


// Converters to promote slices of interfaces, needed when using variadic parameters called using a slice since go doesn't automatically promote a narrow interface inside the slice to be able to use a broader interface.
// for example: without these you couldn't use a slice of LimitedSignal's in a variadic call to a func requiring Signal's. (when you can use separate LimitedSignal's in the same call.)

// converts to []Signal
func PromoteToSignals(s interface{}) []Signal {
	var out []Signal 
	switch st := s.(type) {
	case []Signal:
		return st
	case []LimitedSignal:
		out = make([]Signal, len(st))
		for i := range out {
			out[i] = st[i].(Signal)
		}
	case []PeriodicLimitedSignal:
		out = make([]Signal, len(st))
		for i := range out {
			out[i] = st[i].(Signal)
		}
	case []PeriodicSignal:
		out = make([]Signal, len(st))
		for i := range out {
			out[i] = st[i].(Signal)
		}
	}
	return out
}

// converts to []LimitedSignal
func PromoteToLimitedSignals(s interface{}) []LimitedSignal {
	var out []LimitedSignal 
	switch st := s.(type) {
	case []LimitedSignal:
		return st
	case []PeriodicLimitedSignal:
		out = make([]LimitedSignal, len(st))
		for i := range out {
			out[i] = st[i].(LimitedSignal)
		}
	}
	return out
}

// converts to []PeriodicSignal
func PromoteToPeriodicSignals(s interface{}) []PeriodicSignal {
	var out []PeriodicSignal 
	switch st := s.(type) {
	case []PeriodicSignal:
		return st
	case []PeriodicLimitedSignal:
		out = make([]PeriodicSignal, len(st))
		for i := range out {
			out[i] = st[i].(PeriodicSignal)
		}
	}
	return out
}


