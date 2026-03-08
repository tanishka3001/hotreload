package debounce

import (
	"time"
)

func Debounce(delay time.Duration, input <-chan struct{}) <-chan struct{} {

	output := make(chan struct{})

	go func() {

		var timer *time.Timer

		for range input {

			if timer != nil {
				timer.Stop()
			}

			timer = time.AfterFunc(delay, func() {
				output <- struct{}{}
			})
		}

	}()

	return output
}