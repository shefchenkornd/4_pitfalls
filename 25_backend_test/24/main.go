package main

import "fmt"

// Consider the following code:
func main() {
	fmt.Println("A")

	c := make(chan string)
	c <- "John"

	fmt.Println("B")
}

// Answer: fatal error: all goroutines are asleep - deadlock!