package main

import "fmt"

func main() {
	var p *int // nil

	// fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference

	b := 123
	p = &b

	fmt.Println(*p)

	pf := new(float64)
	fmt.Println(pf)
	fmt.Println(*pf)

	/**
	Глава: Указатели как параметры функции
	По умолчанию все параметры передаются в функцию по значению
	 */
}
