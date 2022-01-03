package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age int
	contactInfo contact
}

type Node struct {
	value int
	next *Node
}

func main() {
	var p person = person{"Ilya", 32, contact{"email@gmail.com", "12345678"}}

	fmt.Println(p.contactInfo.email)
}
