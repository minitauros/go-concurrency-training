package solutions

import (
	"fmt"
	"testing"
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
			expected:    SomeErr,
		},
		{
			description: "Input higher than 1",
			input:       2,
			expected:    nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			res := Sum(tc.input)
			if res != tc.expected {
				t.Errorf("expected: %v; got: %v", tc.expected, res)
			}
		})
	}
}
