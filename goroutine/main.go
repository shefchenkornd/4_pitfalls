package main

import "fmt"

var a = "fsdf"

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
		if !ok {
			break
		}

		fmt.Println(n, ok)
	}
}

func generateNumbers(n int, ch chan int) {
	for i := 0; i <= n; i++ {
		ch <- i * 2
	}

	close(ch)
}
