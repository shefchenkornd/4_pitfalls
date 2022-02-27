package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// What is the output?
func main() {
 var s Shape
 fmt.Println("value of s is", s)
 fmt.Printf("type of s is %T\n", s)
}
// Answer:
// value of s is <nil>
// type of s is <nil>