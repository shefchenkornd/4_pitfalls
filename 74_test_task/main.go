package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(int)
	fmt.Println(s)
}
