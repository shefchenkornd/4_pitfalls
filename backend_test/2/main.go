package main

import "fmt"

// Consider the following code:
// what is the output of the following program?
func main() {
	s := "123"
	ps := &s         // 0xc000010230
	b := []byte(*ps) // [49 50 51]
	pb := &b         // &[49 50 51]

	s += "4"
	*ps += "5"
	b[1] = '0'

	print(*ps) // 12345
	fmt.Println()
	print(string(*pb)) // 103
}
