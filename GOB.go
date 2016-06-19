package signals

import (
	"encoding/gob"
	"io"
)

// write Gob encoding
func Save(p io.Writer,c Signal) error {
	return gob.NewEncoder(p).Encode(&c)
}

// read Gob encoding
func Load(p io.Reader) (s Signal,err error) {
	err=gob.NewDecoder(p).Decode(&s)
	return
}


