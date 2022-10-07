package main

import (
	"fmt"
	"sort"
)

func main() {
	res := MinMax([]int{1, 23, 3, 9, 5})

	fmt.Println(res)
}

// MinMax returns both the minimum and maximum number of the given list/array.
// [1,2,3,4,5] --> [1,5]
func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{
		arr[0],
		arr[len(arr)-1],
	}

	// -OR-
	//
	// min, max := arr[0], arr[0]
	//
	// for _, v := range arr[1:] {
	// 	if min > v {
	// 		min = v
	// 	} else if v > max {
	// 		max = v
	// 	}
	// }
}
