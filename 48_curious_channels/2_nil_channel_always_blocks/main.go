package main

import (
	"fmt"
	"time"
)

// WaitMany In the rewritten WaitMany() we nil the reference to a or b once they have received a value.
// When a nil channel is part of a select statement, it is effectively ignored, so niling a removes it from selection,
// leaving only b which blocks until it is closed, exiting the loop without spinning.
func WaitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
	}
}

func main() {
	chA, chB := make(chan bool), make(chan bool)
	t0 := time.Now()

	go func() {
		close(chA)
		close(chB)
	}()

	WaitMany(chA, chB)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0)) // waited 8.977µs for WaitMany
}
