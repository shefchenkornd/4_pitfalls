package main

import "fmt"

//
func main() {
	ch := make(chan int, 10)

	ch = nil
	fmt.Println(len(ch), cap(ch)) // 0, 0

	ch <- 2 // fatal error: all goroutines are asleep - deadlock!
}
