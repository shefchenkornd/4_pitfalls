package main

import (
	"fmt"
)

// Round up to the next multiple of 5
// input:    output:
// 0    ->   0
// 2    ->   5
// 3    ->   5
// 12   ->   15
// 21   ->   25
// 30   ->   30
// -2   ->   0
// -5   ->   -5
// etc.
func main() {
	res := RoundToNext5(-7)
	fmt.Println("Result = ", res)
	fmt.Println(-2/5)
}

//RoundToNext5 round up to the next multiple of 5
func RoundToNext5(n int) int {
	if (n % 5) == 0 {
		return n
	}

	if n > 0 {
		return ((n/5) + 1) * 5
	} else {
		return (n/5) * 5
	}
}
