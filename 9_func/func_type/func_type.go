package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func action(a, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Println(result)
}

func main() {
	action( 3, 4, add)
	action( 15, 20, sub)
}
