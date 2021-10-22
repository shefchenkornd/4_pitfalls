package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	time.Sleep(3 * time.Second)

	select {
	case n := <-ch:
		fmt.Println(n)
	default:
		fmt.Println("Ничего")
	}
}