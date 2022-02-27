package main

import "fmt"

func defFooStart()  {
	fmt.Println("defFooStart() executed")
}

func defFooEnd()  {
	fmt.Println("defFooEnd() executed")
}

func defMain()  {
	fmt.Println("defMain() executed")
}

func foo()  {
	fmt.Println("foo() executed")
	defer defFooStart()
	panic("panic from foo()")
	defer defFooEnd()
	fmt.Println("foo() done")
}

// Consider the following code:
// What is the correct order of the printed statements?
func main() {
	fmt.Println("main() started")
	defer defMain()
	foo()
	fmt.Println("main() done")
}

// Answer:
// main() started | foo() executed | defFooStart() executed | defMain() executed | panic from foo()
//
// Функции с defer выполняются даже при появлении ошибки.
// Как работает panic:
// 1. Выполнение функции прекращается.
// 2. Любая функция, вызываемая ключевым словом defer, выполняется.
// 3. Выполнение программы завершается.