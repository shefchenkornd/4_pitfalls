package main

import "fmt"

// что выведен данная функция?
func main() {
	fmt.Println("res", test())
}

func test() (x int) {
	defer func(n int) {
		fmt.Println("defer 1func", n)
		fmt.Println("defer 2func", x)

		x *= 10
	}(x)


	x = 1
	return 9
}