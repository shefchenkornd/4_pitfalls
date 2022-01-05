package main

import "sync"

// Неправильно
type IncorrectLock sync.Mutex

// Правильно
type MyLock struct {
	sync.Mutex
}

// Если вам нужны методы из исходного типа, вы можете задать новый тип структуры, встроив исходный в качестве анонимного поля.
func main() {
	var i IncorrectLock
	i.Lock() // ошибка


	var lock MyLock
	lock.Lock() // ok
	lock.Unlock()
}
