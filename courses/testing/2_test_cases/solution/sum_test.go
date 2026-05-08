package solution

import (
	"testing"

	"github.com/minitauros/go-concurrency-training/courses/testing/2_test_cases"
)

func Test_Sum(t *testing.T) {
	t.Run("If input is lower than 2, returns error", func(t *testing.T) {
		err := test_cases.Sum(1)
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("If input is equal to 2, returns nil", func(t *testing.T) {
		err := test_cases.Sum(2)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})

	t.Run("If input is higher than 2, returns nil", func(t *testing.T) {
		err := test_cases.Sum(3)
		if err != nil {
			t.Error("expected nil, but got error")
		}
	})
}
