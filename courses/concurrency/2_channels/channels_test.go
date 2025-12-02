package channels

import (
	"reflect"
	"testing"
	"time"
)

func TestChannels1(t *testing.T) {
	expectedVal := "foo"
	ch := make(chan string, 1)
	ch <- expectedVal

	resCh := make(chan string)
	go func() {
		val := channels1(ch)
		resCh <- val
	}()

	select {
	case <-time.After(20 * time.Millisecond):
		t.Error("did not receive a result within a reasonable time")
	case val := <-resCh:
		if val != expectedVal {
			t.Errorf("got %s, expected %s", val, expectedVal)
		}
	}
}

func TestChannels2(t *testing.T) {
	input := []string{"foo", "bar", "baz"}
	inputCh := make(chan string, len(input))
	outputCh := make(chan string, len(input))

	for _, s := range input {
		inputCh <- s
	}
	close(inputCh)

	go channels2(inputCh, outputCh)

	res := make([]string, 0, len(input))
	go func() {
		for s := range outputCh {
			res = append(res, s)
		}
	}()

	time.Sleep(50 * time.Millisecond)
	if !reflect.DeepEqual(input, res) {
		t.Errorf("got %v, expected %v", res, input)
	}
}

func TestChannels3(t *testing.T) {
	input := []string{"foo", "bar", "baz"}
	outputCh := channels3(input)

	res := make([]string, 0, len(input))
	go func() {
		for val := range outputCh {
			res = append(res, val)
		}
	}()

	time.Sleep(50 * time.Millisecond)
	if !reflect.DeepEqual(input, res) {
		t.Errorf("got %v, expected %v", res, input)
	}
}
