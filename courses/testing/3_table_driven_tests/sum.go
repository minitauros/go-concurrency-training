package tabledriven

import (
	"errors"
)

var SomeErr = errors.New("foo")

func Sum(val int) error {
	if val >= 2 {
		return nil
	}
	return SomeErr
}
