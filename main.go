package main

import (
	"fmt"
)

func main() {
	res := FindOdd([]int{1,1,2})
	fmt.Println("Result = ", res)
}

func FindOdd(seq []int) int {
	m :=  make(map[int]int, len(seq))
	for _, v := range seq {
		m[v] += 1
	}

	for key, value := range m {
		if value % 2 == 1 {
			return key
		}
	}

	//fmt.Println(m)
	return 0
}
