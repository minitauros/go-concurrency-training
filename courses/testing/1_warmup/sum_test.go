package main

import (
	"testing"
)

func Test_sum(t *testing.T) {
	err := Sum(3)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
