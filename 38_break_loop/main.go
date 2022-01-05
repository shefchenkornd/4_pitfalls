package main

import "fmt"

// Выражение break без метки (label) выводит вас только из внутреннего блока switch/select.
// Если использовать выражение return — не вариант, тогда лучший выход — задать метку для внешнего цикла.
func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}

	fmt.Println("out!")
}