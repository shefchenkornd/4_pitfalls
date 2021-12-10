package main

import "fmt"

// What will be the output of the following code?
func main() {
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a))
}
// Answer: [] 0 0
// т.е. присвоив nil мы самым сделали remove all elements