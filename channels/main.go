package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339)) // 2021-10-17 21:19:34.038009831 +0300 MSK m=+0.000030233

	result1 := make(chan int)
	result2 := make(chan int)

	// очень сложные вычисления
	go calculateSomething(result1, 1000)

	// а эти еще сложнее
	go calculateSomething(result2, 2000)

	fmt.Println(<-result1)
	fmt.Println(<-result2)
	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}

func calculateSomething(c chan int, n int) {
	t := time.Now()

	result := 0
	for i := 0; i < n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Результат: %d; Прошло времени: %s\n", result, time.Since(t))
	c <- result
}
