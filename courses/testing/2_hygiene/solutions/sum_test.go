package solutions

import (
	"testing"

	hygiene "github.com/minitauros/go-concurrency-training/courses/testing/2_hygiene"
)

func Test_sum(t *testing.T) {
	t.Run("If input is lower than 2, returns error", func(t *testing.T) {
		err := hygiene.Sum(1)
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("If input is equal to 2, returns nil", func(t *testing.T) {
		err := hygiene.Sum(2)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})

	t.Run("If input is higher than 2, returns nil", func(t *testing.T) {
		err := hygiene.Sum(3)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})
}
