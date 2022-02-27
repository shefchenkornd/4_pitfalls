package main

import "fmt"

/*
Исходный код. Надо исправить две ошибки на линии A and B
func main() {
	var m map[string]int // A
	m["a"] = 1
	if v := m["b"]; v != nil { // B
		println(v)
	}
}
*/

// Правильный ответ
func main() {
	var m = make(map[string]int) // надо использовать make()
	m["a"] = 1

	if v, ok := m["a"]; ok{ // v,ok := map["a"]
		fmt.Println(v)
	}
}
// Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn’t point to an initialized map.