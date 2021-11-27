package main

import (
	"fmt"
	"os"
)

func main() {
	// Стандартным потоком вывода в Go является объект os.Stdout, который фактически представляет консоль
	fmt.Fprintln(os.Stdout, 123) // analog fmt.Println(123)

	// В Go имеется объект os.Stdin, который реализует интерфейс io.Reader и позволяет считывать данные с консоли.
	var name string
	fmt.Fscanln(os.Stdin, &name)
	fmt.Println(name)
}
