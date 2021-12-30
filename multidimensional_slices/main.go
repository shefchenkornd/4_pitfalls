package main

import "fmt"

func main() {
	table := make([][]int, 3) // len(table)=3

	for i := range table {
		table[i] = make([]int, 5)
	}
	fmt.Println(table)
}
