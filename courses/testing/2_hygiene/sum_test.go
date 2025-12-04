package hygiene

import (
	"testing"
)

func Test_sum(t *testing.T) {
	t.Run("If sum is lower than 2, returns error", func(t *testing.T) {
		err := Sum(1)
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("If sum is higher than 2, returns nil", func(t *testing.T) {
		err := Sum(3)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})

	t.Run("If sum is equal to 2, returns nil", func(t *testing.T) {
		err := Sum(2)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})
}
