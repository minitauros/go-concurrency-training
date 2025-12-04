package main

import (
	"errors"
)

func Sum(val int) error {
	if val >= 2 {
		return nil
	}
	return errors.New("foo")
}
