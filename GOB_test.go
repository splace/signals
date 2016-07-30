package signals

import (
	"fmt"
	"testing"
)

func TestGOBSaveLoadTone(t *testing.T) {
	m := Sine{unitX/1000}
	err := SaveGOB("./test output/tone",m)
	if err != nil { t.Error(err)}

	s,err := LoadGOB("./test output/tone")
	if err != nil { t.Error(err)}

	if fmt.Sprintf("%#v", s) != fmt.Sprintf("%#v", m) {
		t.Errorf("%#v != %#v", s, m)
	}

}


func TestGOBSaveLoadStack(t *testing.T) {
	m := Stacked{Sine{unitX/450},Sine{unitX/350}}

	err := SaveGOB("./test output/stack",m)
	if err != nil { t.Error(err)}

	s,err := LoadGOB("./test output/stack")
	if err != nil { t.Error(err)}

	if fmt.Sprintf("%#v", s) != fmt.Sprintf("%#v", m) {
		t.Errorf("%#v != %#v", s, m)
	}

}



