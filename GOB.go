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
func Load(p io.Reader) (Signal,error) {
	var c Signal
	err:=gob.NewDecoder(p).Decode(&c)
	return c,err
}


