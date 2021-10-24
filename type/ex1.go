package main

import "fmt"

type mile int
type kilometer int

func distanceToEnemy(distance mile)  {
	fmt.Printf("Расстояние до противника: %d миль", distance)
}


func main() {
	var distance mile = 7
	distanceToEnemy(distance)

	//var distance2 kilometer = 18
	//distanceToEnemy(distance2)	// ! ошибка, потому что тип не совпадает
	// Таким образом, с помощью именнованных типов мы придаем типу некоторый дополнительный смысл.
}
