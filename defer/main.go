package main

import "fmt"

func main() {
	// Если несколько функций вызываются с оператором defer, то defer ,будет работать по принципу LIFO (англ. Last In, First Out – «последним пришёл — первым ушёл»)
	defer finish()
	defer fmt.Println("Program is working")
	fmt.Println("Program has been started")
}

func finish()  {
	fmt.Println("Program has been finished")
}