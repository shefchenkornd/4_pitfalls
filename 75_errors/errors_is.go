package main

import (
	"errors"
	"fmt"
)

var (
	someErr = errors.New("some error")
)

func do() error {
	return someErr
}

func wrap() error {
	return fmt.Errorf("validateInput: %w", someErr) // Важно: нужно использовать аргумент `%w`
}

// Используя функцию errors.Is(), мы можем обнаружить someErr, даже если она обёрнута, поскольку эта функция проверяет,
// соответствует ли какая-либо ошибка в цепочке обернутых ошибок цели
func main() {
	myErr := someErr
	fmt.Println(myErr == someErr) // true

	// case #1
	err := do()
	fmt.Println(errors.Is(err, someErr)) // true
	fmt.Println(err == someErr)          // true

	// case #2
	err = wrap()
	fmt.Println(errors.Is(err, someErr)) // true
	fmt.Println(err == someErr)          // false
}
