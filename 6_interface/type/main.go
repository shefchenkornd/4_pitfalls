package main

import "fmt"

func main() {
	var a interface{} // реализуется ЛЮБЫМ типом!!!

	a = "hello"
	fmt.Printf("%v - %T \n", a, a) // hello - string

	// invalid operation: a + "text" (mismatched types interface{} and string)
	// fmt.Println(a + "text")

	a = 111
	fmt.Printf("%v - %T \n", a, a) // 111 - int
	fmt.Println(len(a.(string))) // panic: interface conversion: interface {} is int, not string

	a = []bool{false, true}
	fmt.Printf("%v - %T \n", a, a) // [false true] - []bool
}

// Ответ:
// 1. пустой interface{} реализуется ЛЮБЫМ типом
// 2. Для работы с переменой (типа interface{}) требуется привести к КОНКРЕТНОМУ типу (type assertion|type switch)
// 3. Неправильная работа приведёт к панике
