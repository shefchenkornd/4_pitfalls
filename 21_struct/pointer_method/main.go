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
	fmt.Println(p.age)


	pointerP := &p
	pointerP.updateAge(30)
	fmt.Println(pointerP.age)
}
