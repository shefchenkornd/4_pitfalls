package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{}
type Airplane struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Airplane) move() {
	fmt.Println("Самолёт летит")
}

func drive(v Vehicle)  {
	v.move()
}

// Полиморфизм на примере имплементации метода интерфейса и структур
func main() {
	vesta := Car{}
	vesta.move()

	boing := Airplane{}
	boing.move()

	// OR
	var tesla Vehicle = Car{}
	tesla.move()

	var tu104 Vehicle = Airplane{}
	tu104.move()

	// OR
	drive(vesta)
	drive(boing)
}
