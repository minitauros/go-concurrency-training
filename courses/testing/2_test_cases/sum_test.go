package hygiene

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	err := Sum(3)
	if err == nil {
		t.Error("expected error, but got nil")
	}
}
