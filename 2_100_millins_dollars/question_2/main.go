package main

import "fmt"

type myError struct {
	code int
}

func (e myError) Error() string {
	return fmt.Sprintf("code: %d", e.code)
}

func run() error {
	var e *myError
	if false {
		e = &myError{code: 123}
	}
	return e
}

// Вопрос по Golang на миллион долларов
func main() {
	err := run() // nil
	if err != nil {
		fmt.Println("failed to run, error", err) // this line code run
	} else {
		fmt.Println("success")
	}
}

// если на #10 поменять строку:
// var e error
// то в консоли выведется:
// >> success
