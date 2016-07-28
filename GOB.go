package signals

import (
	"encoding/gob"
	"io"
	"os"
)

// write Gob encoding
func WriteGOB(p io.Writer,c Signal) error {
	return gob.NewEncoder(p).Encode(&c)
}

// read Gob encoding
func ReadGOB(p io.Reader) (s Signal,err error) {
	err=gob.NewDecoder(p).Decode(&s)
	return
}

// save Gob encoding
func SaveGOB(pathTo string,s Signal) error {
	file, err := os.Create(pathTo+".gob")
	if err != nil {return err}
	err=gob.NewEncoder(file).Encode(&s)
	if err != nil {return err}
	return file.Close()
}

// load Gob encoding
func LoadGOB(pathTo string) (s Signal,err error) {
	file, err := os.Open(pathTo+".gob")
	if err != nil {return nil,err}
	err=gob.NewDecoder(file).Decode(&s)
	if err != nil {return nil,err}
	err=file.Close()
	return
}


