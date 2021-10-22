package main

import "fmt"

func main() {
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}
