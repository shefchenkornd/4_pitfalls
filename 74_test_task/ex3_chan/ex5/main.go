package main

import (
	"fmt"
	"sync"
)

// Буферизированный канал и select{}
func main() {
	b := make(chan int, 3)
	wg := sync.WaitGroup{}

	// b <- 1 // поиграйся с этим (раскомментируй/закомментируй)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case a := <-b: 			// если в буфере канала, что-то есть, то выполнился этот case:
				fmt.Println(a)
			default: 				// если нечего читать, то выполнился default
				fmt.Println("default")
			}
			break
		}


	}()
	wg.Wait()
}
