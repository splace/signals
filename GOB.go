package signals

import (
	"encoding/gob"
	"io"
	"os"
)

// write Gob encoding
func WriteGOB(p io.Writer,s Signal) error {
	return gob.NewEncoder(p).Encode(&s)
}

// read Gob encoding
func ReadGOB(p io.Reader,s *Signal) error {
	return gob.NewDecoder(p).Decode(s)
}

// save Gob encoding
func SaveGOB(pathTo string,s Signal) error {
	file, err := os.Create(pathTo+".gob")
	if err != nil {return err}
	err=WriteGOB(file,s)
	if err != nil {return err}
	return file.Close()
}

// load Gob encoding
func LoadGOB(pathTo string) (s Signal,err error) {
	file, err := os.Open(pathTo+".gob")
	if err != nil {return nil,err}
	err=ReadGOB(file,&s)
	//err=gob.NewDecoder(file).Decode(&s)
	if err != nil {return nil,err}
	err=file.Close()
	return
}


