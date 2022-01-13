package main

import "fmt"

// This post explores four of the less common properties of channels:
// * A send to a nil channel blocks forever
// * A receive from a nil channel blocks forever
// * A send to a closed channel panics
// * A receive from a closed channel returns the zero value immediately

// A receive from a closed channel returns the zero value immediately
func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", <-c) // prints 1 2 3 0 0
	}

	// The correct solution to this problem is to use a for range loop.
	// for v := range c {
	// 	fmt.Printf("%d ", v) // prints 1 2 3
	// }
	// or
	// for v, ok := <- c; ok ; v, ok = <- c {
	// 	fmt.Printf("%d ", v) // prints 1 2 3
	// }
}

// Once a channel is closed, and all values drained from its buffer, the channel will always return zero values immediately.
