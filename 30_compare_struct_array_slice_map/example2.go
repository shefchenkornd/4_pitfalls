package main

import "fmt"

type data2 struct {
	num    int               // ok
	checks [10]func() bool   // не сравниваемо
	doit   func() bool       // не сравниваемо
	m      map[string]string // не сравниваемо
	bytes  []byte            // не сравниваемо
}

// Если хоть одно из полей не сравниваемо, то применение оператора эквивалентности приведёт к ошибке компилирования.
// Обратите внимание, что сравнивать массивы можно только тогда, когда сравниваемы их данные.
func main() {
	v1 := data2{}
	v2 := data2{}
	fmt.Println("v1 == v2:", v1 == v2) // invalid operation: v1 == v2 (struct containing [10]func() bool cannot be compared)
}
