package main

import "fmt"

type MyInterface interface {}

type Example struct {
	Value string
}

func example1() MyInterface {
	var e *Example // nil
	return e
}

func example2() MyInterface {
	return nil
}

func main() {
	// Вопрос по Golang на миллион долларов
	fmt.Println(example1() == example2()) // Ответ: false
	// Объяснение:
	// значение у обоих функций будет рано nil, а вот типы будут отличаться *main.Example, <nil>

	// fmt.Println(example1(), example2())
	// fmt.Printf("%T, %T", example1(), example2())
}
