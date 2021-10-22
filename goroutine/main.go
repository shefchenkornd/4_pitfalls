package main

import "fmt"

func main() {
	ch := make(chan int)

	go generateNumbers(5, ch)

	// 1й способ чтения из канала
	//for n := range ch {
	//	fmt.Println(n)
	//}

	// 2й способ чтения из канала
	for {
		n, ok := <-ch
		fmt.Println(n, ok)

		if !ok {
			break
		}
	}
}

func generateNumbers(n int, ch chan int) {
	for i := 0; i <= n; i++ {
		ch <- i * 2
	}

	close(ch)
}
