package main

import "fmt"

func makeSquares(array [10]int) {
	for index, elem := range array {
		array[index] = elem * elem
	}
}

// Consider the following code:
// If we pass the array to makeSquares(), what is the output?
func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	makeSquares(a)

	fmt.Println(a) // [0 1 2 3 4 5 6 7 8 9]
}
// Answer: массив не передаётся по ссылке, поэтому makeSquares(a) не изменит наш массив, и он распечатается в исходном виде
