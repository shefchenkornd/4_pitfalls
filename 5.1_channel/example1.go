package main

import "fmt"

func main() {
	intCh := make(chan int)

	for i := 0; i <= 5; i++ {
		go factorial(i, intCh)

		fmt.Printf("i = %d, factorial = %d \n", i, <-intCh)
	}
}

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result // отправка данных в канал
}
