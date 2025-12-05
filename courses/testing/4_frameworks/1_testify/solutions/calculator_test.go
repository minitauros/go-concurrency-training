package solutions

import (
	"fmt"
	"testing"

	frameworks "github.com/minitauros/go-concurrency-training/courses/testing/4_frameworks"
	"github.com/stretchr/testify/assert"
)

func Test_Calculator_Sum(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		description string
		inputs      []int64
		expected    int64
	}{
		{
			description: "Works with numbers above zero",
			inputs:      []int64{1, 2, 5},
			expected:    8,
		},
		{
			description: "Works with numbers below zero",
			inputs:      []int64{-1, -2, -4},
			expected:    -7,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			calculator := &frameworks.Calculator{}
			res := calculator.Sum(tc.inputs...)
			assert.Equal(tc.expected, res)
		})
	}
}

func Test_Calculator_Multiply(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		description string
		start       int64
		multiplyBy  int64
		expected    int64
	}{
		{
			description: "Works with numbers above zero",
			start:       10,
			multiplyBy:  10,
			expected:    100,
		},
		{
			description: "Works with numbers below zero",
			start:       -10,
			multiplyBy:  -8,
			expected:    -80,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
			calculator := &frameworks.Calculator{}
			res := calculator.Multiply(tc.start, tc.multiplyBy)
			assert.Equal(tc.expected, res)
		})
	}
}

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
			calculator := &frameworks.Calculator{}
			res := calculator.SpecialSub(tc.start, tc.sub)
			assert.Equal(tc.expected, res)
		})
	}
}
