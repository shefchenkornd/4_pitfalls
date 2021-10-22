package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(x, y int) int {
	return x - y
}

func action(a, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Println(result)
}

func main() {
	action( 3, 4, add)
}
