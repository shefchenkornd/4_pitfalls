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
	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	ch1 := make(chan int, 10000)
	ch2 := make(chan int, 10000)
	ch3 := make(chan int, 10000)

	for i := 0; i < 10000; i++ {
		ch1 <- i
	}
	close(ch1)

	for i := 0; i < 10000; i++ {
		ch2 <- i
	}
	close(ch2)

	for i := 0; i < 10000; i++ {
		ch3 <- i
	}
	close(ch3)

	mergeChan := merge(ch1, ch2)
	for val := range mergeChan {
		fmt.Println("Val =", val)
	}

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println(time.Since(start))
}

// результирующий канал С БУФЕРОМ
func merge(cs ...<-chan int) <-chan int {
	l := 0
	for _, c := range cs {
		l += cap(c)
	}

	resultCh := make(chan int, l)

	wg := sync.WaitGroup{}
	for _, c := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()

			for v := range ch {
				resultCh <- v
			}
			return
		}(c)
	}
	wg.Wait()
	close(resultCh)

	return resultCh
}
