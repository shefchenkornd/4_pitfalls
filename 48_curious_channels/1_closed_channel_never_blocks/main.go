package main

import (
	"fmt"
	"sync"
	"time"
)

// So what is going on here? As soon as the finish channel is closed, it becomes ready to receive.
// As all the goroutines are waiting to receive either from their time.After channel,
// or finish, the select statement is now complete and the goroutines exits after calling done.Done() to deincrement the WaitGroup counter.
// This powerful idiom allows you to use a channel to send a signal to an unknown number of goroutines,
// without having to know anything about them, or worrying about deadlock.
func main() {
	// finish := make(chan struct{}) // As the behaviour of the close(finish) relies on signalling the close of the channel, not the value sent or received, declaring finish to be of type chan struct{} says that the channel contains no value; we’re only interested in its closed property.
	var finish chan struct{}
	var done sync.WaitGroup

	for i := 0; i < 100; i++ {
		done.Add(1)
		go func() {
			select {
			case <-time.After(1 * time.Hour):
			case <-finish:
			}
			done.Done()
		}()
	}

	t0 := time.Now()

	close(finish)

	done.Wait() // wait for the goroutine to stop
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0)) // Waited 57.244µs for goroutine to stop
}
