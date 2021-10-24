package main

import "fmt"

type library []string //  Данный тип будет представлять своего рода библиотеку

func printBooks(lib library)  { // При этом функция printBooks работает именно с типом library, а не с любым срезом из строк.
	for _, book := range lib {
		fmt.Println(book)
	}
}

func main() {
	lib := library{"Book1", "Book2", "Book3"}
	printBooks(lib)
}
