package signals

const cacheSize = 256

// a Signal that stores, some, property values, rather than always getting them from the embedded Signal.
type Cached struct {
	Signal
	cache map[x] y
}

func NewCached(s Signal) Cached {
	return Cached{s,make(map[x]y)}
}

func (s Cached) property(offset x) y {
	if v,ok:=s.cache[offset];ok {return v}
	if len(s.cache)>cacheSize+10{
		for i:=range(s.cache){
			delete(s.cache,i)
			if len(s.cache)<=cacheSize{break}
		}
	}
	v:=s.Signal.property(offset)
	s.cache[offset]=v
	return v
}



