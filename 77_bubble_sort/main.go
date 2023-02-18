package main

import "fmt"

func main() {
	sl := []int{3, 5, 1, 22, 0, 103, 5, -1, -2, 77, 505}
	fmt.Printf("unordered slice: %v \n", sl)

	res := bubbleSort(sl)
	fmt.Println("ordered slice:\t", res)
}

func bubbleSort(sl []int) []int {
	for i := 0; i < len(sl); i++ {
		for y := 0; y < len(sl)-1; y++ {
			if sl[y] > sl[y+1] {
				sl[y], sl[y+1] = sl[y+1], sl[y]
			}
		}
	}

	return sl
}
