package debounce

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {

	input := make(chan struct{}, 10)

	output := Debounce(100*time.Millisecond, input)

	for i := 0; i < 5; i++ {
		input <- struct{}{}
	}

	select {
	case <-output:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("debounce did not trigger")
	}
}