package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Hello from service1"
}

func service2(c chan string) {
	c <- "Hello from service2"
}

// Consider the following code:
// After executing this code, what is the output?
func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	fmt.Println("start....")

	select {
	case res := <-chan1:
		fmt.Printf("Response from service №1: \"%v\" \n", res)
	case res := <-chan2:
		fmt.Printf("Response from service №2: \"%v\" \n", res)
	// default:
	// 	fmt.Println("No response received") // выведется вот эта строка!
	}

	fmt.Println("finish....")
}

// Answer: выведется следующий текст "No response received", тк, если бы не было default, то был либо ответ 1 либо 2.
// default в select{} делает его неблокирующим поток main
