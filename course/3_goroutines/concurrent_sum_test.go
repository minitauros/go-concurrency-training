package concurrency

import (
	"testing"
)

func TestConcurrentSum(t *testing.T) {
	input := make([]int, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		input = append(input, i)
	}
	res := sumConcurrent(input)
	expected := 499999500000
	if res != expected {
		t.Errorf("expected: %d, actual: %d", expected, res)
	}
}
