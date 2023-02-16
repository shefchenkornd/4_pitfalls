package main

import (
	"bytes"
	"fmt"
	"reflect"
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
	fmt.Println("v1 == v2:", v1 == v2) // выводит: v1 == v2: true

	// Compare slice of byte.

	sliceByte1 := []byte{'G', 'E', 'E', 'K', 'S'}
	sliceByte2 := []byte{'G', 'E', 'E', 'K', 'S'}
	// Сравнивать slice of byte напрямую нельзя!
	// fmt.Println(sliceByte1 == sliceByte2) // Error: "Invalid operation: sliceByte1 == sliceByte2 (the operator == is not defined on []byte)"

	// Using Compare function
	res := bytes.Compare(sliceByte1, sliceByte2)
	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!") // this is output
	}

	// либо же при помощи
	fmt.Println("Слайсы равны? ", reflect.DeepEqual(sliceByte1, sliceByte2))
}
