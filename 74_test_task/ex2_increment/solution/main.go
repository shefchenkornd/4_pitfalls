package main

import "fmt"

type Count int

func (c *Count) Increment() {
	*c++
}

// Задача: нужно исправить метод Increment(), чтобы он выполнял инкремент
func main() {
	var count Count
	count.Increment()
	fmt.Print(count)
}
