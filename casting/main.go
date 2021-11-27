package main

import "fmt"

func main() {
	var i int = 10
	var a int32= 10
	fmt.Printf("%v - %T \n", i, i)
	fmt.Printf("%v - %T \n", a, a)

	var c int32
	c = int32(i)

	fmt.Println(c == a)
}
