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
	fmt.Println(example1() == example2()) // false
}
