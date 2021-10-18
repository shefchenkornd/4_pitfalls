package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339)) // 2021-10-17 21:19:34.038009831 +0300 MSK m=+0.000030233

	// очень сложные вычисления
	go calculateSomething(1000)

	// а эти еще сложнее
	go calculateSomething(2000)

	time.Sleep(8 * time.Second)
	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}

func calculateSomething(n int) {
	t := time.Now()

	result := 0
	for i := 0; i < n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Результат: %d; Прошло времени: %s\n", result, time.Since(t))
}
