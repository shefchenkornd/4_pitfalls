package main

import "fmt"

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40, 50} { // <-- переменная value создана на время выполнения цикла
		// на каждой новой итерации этот value переиспользуется, иначе было бы расточительно
		// i := value // Решение этой проблемы ниже
		numbers = append(numbers, &value)
	}

	fmt.Println(numbers)

	for _, number := range numbers {
		fmt.Printf("%d ", *number) // 50 50 50 50 50 <- это неправильно
		// 10, 20, 30, 40, 50 <- это ПРАВИЛЬНО, для этого нужно применить решение выше!
	}
}
