package main

// This post explores four of the less common properties of channels:
// * A send to a nil channel blocks forever
// * A receive from a nil channel blocks forever
// * A send to a closed channel panics
// * A receive from a closed channel returns the zero value immediately

// A send to a nil channel blocks forever
func main() {
	var ch chan string
 	ch <- "let's get started" // fatal error: all goroutines are asleep - deadlock!
}
// This example program will deadlock on line 12 because the zero value for an uninitalised channel is nil.