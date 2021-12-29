package main

import "fmt"

type MyStruct struct {
	Name string
}

//func NewMyStruct(name string) *MyStruct {
//	return &MyStruct{Name: name}
//}

func (ms MyStruct) PrintNameMethod() {
	fmt.Println(ms.Name)
}

/**
Дело в том, что функция UpdateNameMethod получила копию MyStruct, которая затем была обновлена.
Но исходный MyStruct остался неизменным.

Надо обновить получатель, чтобы он получал указатель на структуру:
*/

//UpdateNameMethod ...
func (ms *MyStruct) UpdateNameMethod(newName string) {
	ms.Name = newName
}

func main() {
	ms := MyStruct{
		Name: "Joe",
	}

	// Это небольшое изменение дает возможность менять структуру напрямую.
	// Так исходный MyStruct меняется и код выводит то, что от него ожидается.
	ms.UpdateNameMethod("Jane")
	ms.PrintNameMethod()
}
