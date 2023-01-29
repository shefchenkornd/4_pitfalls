package main

import "fmt"

func main() {
	fmt.Println("Main function starts")
	intCh := make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		intCh<-5
	}()

	fmt.Println("Get value from channel", <-intCh) // получение данных из канала
	fmt.Println("The End")
}
