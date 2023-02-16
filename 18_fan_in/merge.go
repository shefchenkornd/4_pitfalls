package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Написать код функции, которая делает merge N каналов. Весь входной поток перенаправляется в один канал.
func main() {
	start := time.Now()
	fmt.Println(runtime.NumGoroutine())


	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)
	for val := range ch3 {
		fmt.Println(val)
	}

	fmt.Println(time.Since(start))
	fmt.Println("Кол-во горутин:", runtime.NumGoroutine())
}

// результирующий канал без буфера
func merge[T any](cs ...chan T) chan T {
	resultCh := make(chan T)
	wg := sync.WaitGroup{}

	for _, singleChan := range cs {
		wg.Add(1)
		singleChan := singleChan

		go func() {
			defer wg.Done()

			for val := range singleChan {
				resultCh <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}
