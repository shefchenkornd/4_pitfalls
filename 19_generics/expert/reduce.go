package main

import "fmt"

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accumulator(init, el)
	}

	return init
}

func main() {
	list := []int{1, 2, 4, 8}

	sum := func(x, y int) int { return x + y }
	mul := func(x, y int) int { return x * y }

	res := Reduce(list, sum, 0)
	fmt.Println(res)

	res2 := Reduce(list, mul, 1)
	fmt.Println(res2)
}
