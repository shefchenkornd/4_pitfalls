package main

import "fmt"

type S struct{}

func f(x interface{}) {}

func g(x *interface{}) {
	fmt.Printf("%T",  x)
}

// Consider the following lines of code, marked with an A, B, C and D.
// Which of these have syntax issues (choose one or more)?
func main() {
	s := S{}
	p := &s

	f(s) // A
	g(s) // B
	f(p) // C
	g(p) // D

}
// Answer: B and D