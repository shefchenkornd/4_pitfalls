package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // если не закрыть, то будет deadlock!

	for elem := range queue {
		fmt.Println(elem)
	}

// 	Если НЕ закрыть канал, то мы его прочитаем до конца, а потом будет deadlock!
// 	fatal error: all goroutines are asleep - deadlock!
}
