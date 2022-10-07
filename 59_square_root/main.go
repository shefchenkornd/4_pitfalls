package main

import (
	"fmt"
	"math"
)

// isImperfectSquare if the parameter is itself not a perfect square then -1 should be returned.
// You may assume the parameter is non-negative.
const isImperfectSquare = -1

func main() {
	res := FindNextSquare(121)

	fmt.Println(res)
}

// FindNextSquare
// You might know some pretty large perfect squares. But what about the NEXT one?
// Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter.
// Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.
// If the parameter is itself not a perfect square then -1 should be returned. You may assume the parameter is non-negative
// 121 --> 144
// 625 --> 676
// 114 --> -1 since 114 is not a perfect square
func FindNextSquare(sq int64) int64 {
	x := math.Sqrt(float64(sq))
	res := int64(x)

	if float64(res) != x {
		return isImperfectSquare
	}

	res += 1
	return int64(math.Pow(float64(res), 2))
}
