package main

import "fmt"

func service() {
	fmt.Println("Hello from service")
}

// what is the output of the following program?
func main() {
	fmt.Println("main() started")

	go service()
	select {} // fatal error: all goroutines are asleep - deadlock!
	// Пояснение: But if this empty select is put in the main goroutine then it will cause a deadlock.

	fmt.Println("main() stopped")
}
