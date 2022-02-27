package main

import "fmt"

// Consider the following code:
// After executing this code, what is the output?
func main() {
	a := [...]int{1, 2, 3, 4, 5,6, 7, 8, 9} // len=7, cap=7
	s := a[2:4] // 3, 4

	newS := append(s, 55, 66) // 3, 4, 55, 66

	fmt.Println(newS)
	fmt.Println(len(newS), cap(newS)) // 4, 7
}

// Answer: cap() не поменялось, потому что мы не вышли за него,
// поэтому у переменной newS len=4, cap=7