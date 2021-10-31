package main

import "fmt"

type library []string

type I interface {

}

func (l library) print()  {
	for _, book := range l {
		fmt.Println(book)
	}
}

func main() {
	l := library{"Book1", "Book2", "Book3"}
	l.print()
}
