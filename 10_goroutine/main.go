package main

import "fmt"

var a = "fsdf"

func main() {
	ch := make(chan int)

	for i := 1; i < 10; i++{
		for j := 1; j < 10; j++{
			fmt.Print(i * j, "\t")
		}
		fmt.Println()
	}

	go generateNumbers(5, ch)

	// 1-й способ чтения из канала
	//for n := range ch {
	//	fmt.Println(n)
	//}

	// 2-й способ чтения из канала
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
