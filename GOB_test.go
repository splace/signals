package signals

import (
	"fmt"
	"os"
	"testing"
)

func TestGOBSaveLoadTone(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Create("./test output/tone.gob"); err != nil {panic(err)}else{defer file.Close()}
	m := Sine{unitX/1000}
	if err := Save(file,m); err != nil {
		panic(err)
	}
	file.Close()

	if file, err = os.Open("./test output/tone.gob"); err != nil {
		panic(err)
	}
	defer file.Close()

	s,err := Load(file)
	if err != nil {
		panic(err)
	}
	if fmt.Sprintf("%#v", s) != fmt.Sprintf("%#v", m) {
		t.Errorf("%#v != %#v", s, m)
	}

}

