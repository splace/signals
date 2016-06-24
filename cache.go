package signals

const cacheSize = 256

// a Signal that stores and reuses, some, recent property values, rather than always getting them from the embedded Signal.
type Cached struct {
	Signal
	cache map[x] y
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

// a Signal that stores and reuses, sequential and evenly spaced, recent property values, rather than always getting them from the embedded Signal.
//type Buffered struct {
//	Shifted
//	reader io.Reader
//}
//
//func NewBuffered(s LimitedSignal,sampleBytes uint8,sampleRate uint32) Buffered {
//	r, w := io.Pipe()
//	Encode(w, sampleBytes, sampleRate, s.MaxX(), s)
//	return Buffered{Shifted{s,0},r}
//}
//

