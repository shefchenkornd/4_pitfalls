package main

import (
	"bytes"
	"fmt"
)

type data struct {
	num     int
	fp      float64
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

// Можно использовать оператор эквивалентности == для сравнения переменных структур,
// если каждое поле структуры можно сравнить с помощью этого оператора.
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2) // выводит: v1 == v2: true


	// Compare slice of byte.

	sliceByte1 := []byte{'G', 'E', 'E', 'K', 'S'}
	sliceByte2 := []byte{'G', 'E', 'e', 'K', 'S'}
	// fmt.Println(slice_byte_1 == slice_byte_2) 	// Сравнивать slice of byte напрямую нельзя!

	// Using Compare function
	res := bytes.Compare(sliceByte1, sliceByte2)

	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!") // this is output
	}


}
