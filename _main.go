package _main

import "fmt"

func main() {
	recursion(5, 1, []int64{})
}

func recursion(left, last int64, arr []int64) {
	if left == 0 {
		for _, i := range arr {
			fmt.Println(i)
		}
		fmt.Println()
		return
	}

	b := []int64{}
	for n := last; n < left; n++ {
		b = arr

		b = append(b, n)
		recursion(left-n, n, b)
	}
}
