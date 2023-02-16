package main

import (
	"bytes"
	"fmt"
	"reflect"
)

// Consider the following code:
// Question: what is the output?
func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{0, 1, 2, 3}

	// нельзя сравнивать два slices
	// fmt.Println(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)

	// to compare two slices
	fmt.Println(reflect.DeepEqual(s1, s2))


	// если бы слайсы были бы из byte, то можно было использовать это:
	byte1 := []byte{'a', 'b', 'c'}
	byte2 := []byte{'a', 'b', 'c'}

	res := bytes.Compare(byte1, byte2)
	if res == 0 {
		fmt.Println("Слайсы равны")
	} else {
		fmt.Println("Слайсы НЕ равны")
	}
}
