package main

import "fmt"

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
}
