package solution

import (
	"errors"
	"fmt"
	"testing"

	table_driven "github.com/minitauros/go-concurrency-training/courses/testing/3_table_driven_tests"
)

func Test_Sum(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    error
	}{
		{
			description: "If input is lower than 2, returns error",
			input:       1,
			expected:    table_driven.SomeErr,
		},
		{
			description: "If input is equal to 2, returns nil",
			input:       2,
			expected:    nil,
		},
		{
			description: "If input is higher than 2, returns nil",
			input:       3,
			expected:    nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			res := table_driven.Sum(tc.input)
			if !errors.Is(res, tc.expected) {
				t.Errorf("expected: %v; got: %v", tc.expected, res)
			}
		})
	}
}
