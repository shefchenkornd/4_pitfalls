package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var p Person;
	p = Person{"John", 32}

	poiterP := &p

	fmt.Println(p, poiterP.name)
}
