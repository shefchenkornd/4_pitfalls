package main

import "fmt"

// Аргументы для вызовов отложенных функций вычисляются тогда же, когда и выражение defer (а не когда на самом деле выполняется функция).
func main() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }()) // result => 2 (not ok if you expected 4)
	i++
}
