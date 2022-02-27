package main

import (
	"fmt"
	"reflect"
)

// Consider the following code:
// Question: what is the output?
func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{0, 1, 2, 3}

	// нельзя сравнивать два slices
	fmt.Println(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)

	// to compare two slices
	fmt.Println(reflect.DeepEqual(s1, s2))
}
