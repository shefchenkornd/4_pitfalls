package main

import "fmt"

type S struct {
	a, b, c string
}

// What is the output?
func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(x == y)
	fmt.Println(x, y)
	fmt.Printf("%T", x)
}

// Answer: false
// Было бы true, если строки ниже были бы без pointer'ов на структуры:
// x := interface{}(&S{"a", "b", "c"})
// y := interface{}(&S{"a", "b", "c"})
