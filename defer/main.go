package main

import "fmt"

func main() {
	// Если несколько функций вызываются с оператором defer, то те функции, которые вызываются раньше, будут выполняться позже всех.
	defer finish()
	defer fmt.Println("Program is working")
	fmt.Println("Program has been started")
}

func finish()  {
	fmt.Println("Program has been finished")
}