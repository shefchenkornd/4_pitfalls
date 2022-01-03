package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {
	model string
}

type Airplane struct {
	model string
}

func (c Car) move() {
	fmt.Println(c.model, "едет")
}

func (a Airplane) move() {
	fmt.Println(a.model, "летит")
}

/**
В данном случае определен массив vehicles, который содержит набор структур, которые соответствуют интерфейсу Vehicle,
то есть объекты Car и Aircraft.
То есть объект Vehicle может принимать различные формы: или структуру Car, или структуру Aircraft.
При переборе массива для каждого объекта вызывается метод move.
И в зависимости от реального типа структуры динамически определяется,
какая именно реализация метода move для какой структуры должна вызываться.
 */
func main() {
	car := Car{model: "Vesta"}
	airplane := Airplane{model: "Ту-150"}

	vehicles := [...]Vehicle{car, airplane}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}
