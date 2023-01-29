package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) updateAge(newAge int)  {
	(*p).age = newAge
}

func main() {
	p := person{"John", 27}

	p.updateAge(28)
	fmt.Printf("%v %d years old \n", p.name, p.age)


	pointerP := &p
	pointerP.updateAge(30)
	fmt.Printf("%v %d years old \n", pointerP.name, pointerP.age)
}
