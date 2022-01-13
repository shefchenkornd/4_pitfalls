package main

import "fmt"

// This post explores four of the less common properties of channels:
// * A send to a nil channel blocks forever
// * A receive from a nil channel blocks forever
// * A send to a closed channel panics
// * A receive from a closed channel returns the zero value immediately

// A receive from a nil channel blocks forever
func main() {
	var ch chan string
 	fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}
// So why does this happen ? Here is one possible explanation
//
// The size of a channel’s buffer is not part of its type declaration, so it must be part of the channel’s value.
// If the channel is not initalised then its buffer size will be zero.
// If the size of the channel’s buffer is zero, then the channel is unbuffered.
// If the channel is unbuffered, then a send will block until another goroutine is ready to receive.
// If the channel is nil then the sender and receiver have no reference to each other; they are both blocked waiting on independent channels and will never unblock.