package kwargs

import (
	"fmt"
)

type Kwargs struct {
	args map[string]interface{}
}

func create() Kwargs {
	k := Kwargs{
		args: make(map[string]interface{}),
	}
	return k
}

func (k Kwargs) putString(arg string, v string) error {
	k.args[arg] = v
	return nil
}

func (k Kwargs) getString(arg string) (string, error) {
	for k, v := range k.args {
		if k == arg {
			switch v.(type) {
			case string:
				return v.(string), nil
			default:
				return "", fmt.Errorf("Argument %s is not a string", arg)
			}
		}
	}
	return "", fmt.Errorf("Argument %s not found", arg)
}

func (k Kwargs) get(arg string) (interface{}, error) {
	for k, v := range k.args {
		if k == arg {
			return v, nil
		}
	}
	return "", fmt.Errorf("Argument %s not found", arg)
}
