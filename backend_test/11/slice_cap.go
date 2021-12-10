package main

import "fmt"

// What will be the output of the following code?
func main() {
	a := []string{"A", "B", "C", "D", "E"}
	a = a[:0]
	fmt.Println(a, len(a), cap(a)) // [] 0 5
}
// Keep allocated memory
