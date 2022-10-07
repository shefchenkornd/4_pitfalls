package main

import (
	"fmt"
)

// summation(8) -> 36
// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8
func main() {
	res := Summation(8)
	fmt.Println(res)
}

// Summation
// Write a program that finds the summation of every number from 1 to num.
// The number will always be a positive integer greater than 0.
// summation(2) -> 3
// 1 + 2
// summation(8) -> 36
// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8
func Summation(n int) int {
	return n * (n + 1) / 2

	// var res int
	// for i := 1; i <= n; i++ {
	// 	res += i
	// }
	// return res
}
