package frameworks_testify

import (
	"fmt"
	"testing"

	frameworks_testify "github.com/minitauros/go-concurrency-training/courses/testing/4_frameworks/1_testify"
	"github.com/stretchr/testify/assert"
)

func Test_Calculator_SpecialSub(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		description string
		start       int64
		sub         int64
		expected    int64
	}{
		{
			description: "Works with above zero input",
			start:       5,
			sub:         3,
			expected:    2,
		},
		{
			description: "Works with negative result",
			start:       5,
			sub:         8,
			expected:    -3,
		},
		{
			description: "Works with negative input",
			start:       -5,
			sub:         -4,
			expected:    -1,
		},
		{
			description: "If value is 100, works as normal",
			start:       100,
			sub:         5,
			expected:    95,
		},
		{
			description: "If input is above 100, subtracts double the given value",
			start:       101,
			sub:         5,
			expected:    91,
		},
		{
			description: "If input is exactly 110, returns the answer to the ultimate question of life, the universe, and everything",
			start:       110,
			sub:         61234,
			expected:    42,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			calculator := &frameworks_testify.Calculator{}
			res := calculator.SpecialSub(tc.start, tc.sub)
			assert.Equal(tc.expected, res)
		})
	}
}
