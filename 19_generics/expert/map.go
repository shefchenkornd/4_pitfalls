package main

import "fmt"

func Map[T comparable](list []T, change func(T) T) []T {
	for indx := range list {
		list[indx] = change(list[indx])
	}

	return list
}

func main() {
	list := []int{2, 4, 6, 8}

	square := func(x int) int { return x * x }

	res := Map(list, square)
	fmt.Println(res)
}
