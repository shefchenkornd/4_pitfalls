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

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res)
	case res := <-chan2:
		fmt.Println("Response from service 2", res)
	default:
		fmt.Println("No response received") // выведется вот эта строка, НО Я ПОКА ЧТО не знаю почему
	}
}
// Answer: выведется следующий текст "No response received"