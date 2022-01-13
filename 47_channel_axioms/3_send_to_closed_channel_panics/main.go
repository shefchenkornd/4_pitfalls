package main

import "fmt"

// This post explores four of the less common properties of channels:
// * A send to a nil channel blocks forever
// * A receive from a nil channel blocks forever
// * A send to a closed channel panics
// * A receive from a closed channel returns the zero value immediately

// A send to a closed channel panics
func main() {
	var c = make(chan int, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for j:=0; j<10; j++ {
				c <- j // panic: send on closed channel
			}
			close(c)
		}()
	}

	for i := range c {
		fmt.Println(i)
	}
}
