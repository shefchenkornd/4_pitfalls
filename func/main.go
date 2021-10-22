package main

import "fmt"

func main() {
	add(2, 3)
	add(2, 3)
}

// Неопределенное количество параметров в функции [ставится многоточие: numbers ...int]
func add(numbers ...int) {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum =", sum)
}
