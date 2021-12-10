package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Consider the following code:
// What is the output?
func main() {
	r := Rect{5.0, 4.0}
	var s Shape = r

	area := s.Area()
	fmt.Println(area) // error
}
// Answer: Compile-time error: Rect does not implement Shape (Area method has pointer receiver)
