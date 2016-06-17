package signals

import (
	"encoding/gob"
	"io"
)

// write Gob encoding
func Save(p io.Writer,c interface{}) error {
	return gob.NewEncoder(p).Encode(&c)
}

// read Gob encoding
func Load(p io.Reader) (interface{},error) {
	var c interface{}
	err:=gob.NewDecoder(p).Decode(&c)
	return c,err
}


