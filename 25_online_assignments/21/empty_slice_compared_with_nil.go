package main

import "fmt"

// Consider the following code:
// After executing this code, what is the output?
func main() {
	var s1 []int
	// s1 := make([]int, 3)
	fmt.Println("len = ", len(s1), " | cap = ", cap(s1))

	s2 := []int{1, 2, 3}
	n1 := copy(s1, s2)

	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2) // 0 [] [1 2 3]
	fmt.Println("s1 == nil", s1 == nil)             // true
	fmt.Println(len(s1), cap(s1))                   // 0 0
}

// При использовании функции copy() необходимо учитывать два случая:
// 1. если длина src больше длины dst, то количество скопированных элементов равно длине dst.
// 2. если длина dst больше длины src, то количество скопированных элементов равно длине src.
