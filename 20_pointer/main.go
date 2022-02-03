package main

import "fmt"

func main() {
	var p *int // nil
	// fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference

	// если у нас есть значение и мы хотим получить его указатель, то мы используем оператор `&`:
	b := 123
	p = &b // Оператор `&` генерирует указатель на свой операнд.

	// если мы хотим получить значение переменной по его указателю, то мы используем оператор `*` и указатель:
	fmt.Println("pointer `p` =", p) // 0xc0000b8000
	fmt.Println("Получаем значение по указателю: *p =", *p)         // 123

	pf := new(float64)
	fmt.Println("pointer pf =", pf)  // 0xc0000b8008
	fmt.Println("Получаем значение по указателю: *pf =", *pf) // 0

	/**
	Указатели как параметры функции:
	По умолчанию все параметры передаются в функцию по значению
	*/
}
