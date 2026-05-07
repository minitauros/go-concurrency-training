package solution

import (
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
			description: "Input equal to 1",
			input:       1,
			expected:    table_driven.SomeErr,
		},
		{
			description: "Input higher than 1",
			input:       2,
			expected:    nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			res := table_driven.Sum(tc.input)
			if res != tc.expected {
				t.Errorf("expected: %v; got: %v", tc.expected, res)
			}
		})
	}
}
