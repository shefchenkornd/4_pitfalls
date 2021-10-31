package main

import "fmt"

// вот так надо запускать, иначе будет ошибка "undefined: factorial"
// go run main.go factorial.go
func main() {
	fmt.Println(factorial(10))
}
