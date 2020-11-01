package kwargs

import (
	"testing"
)

func TestArgs(t *testing.T) {
	k := Kwargs{
		args: map[string]interface{}{
			"my_string":  "this is a string",
			"my_float64": 0.55,
		},
	}
	s, err := k.getString("my_string")
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
	if s != "this is a string" {
		t.Errorf("Argument my_string got the unexpected value %s", s)
	}
}

func TestCreate(t *testing.T) {
	k := create()
	k.putString("my_string", "this is a string")
	s, err := k.getString("my_string")
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}
	if s != "this is a string" {
		t.Errorf("Argument my_string got the unexpected value %s", s)
	}
}
