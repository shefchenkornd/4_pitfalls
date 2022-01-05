package main

import "fmt"

type data struct {
	name string
}

// Если у вас есть таблица, состоящая из структур, то вы не можете обновлять отдельные структурные поля.
func main() {
	m := map[string]data{"x": {name: "John"}}

	// 	Это не работает, потому что элементы таблицы не адресуемы.
	// m["x"].name = "abc" // error: cannot assign to struct field m["x"].name in map

	// Решение №1 - использовать временную переменную.
	tmp := m["x"]
	tmp.name = "abc"
	m["x"] = tmp

	fmt.Println(m) // map[x:{abc}]

	// ***************************************************

	// Решение №2 - использовать хеш-таблицу с указателями.
	m2 := map[string]*data{"z": {name: "Mike"}}
	m2["z"].name = "Rob" // ok
	// m2["a"].name = "what?" // panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(m2["z"]) // &{Rob}
}
