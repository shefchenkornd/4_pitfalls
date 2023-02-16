package main

import (
	"fmt"
	"time"
)

// Оператор select позволяет go-процедуре находиться в ожидании нескольких операций передачи данных.
//
// Select блокируется до тех пор, пока один из его блоков case не будет готов к запуску, а затем выполняет этот блок.
// Если сразу несколько блоков могут быть запущены, то выбирается произвольный.
func main() {
	var c chan bool
	c <- true
	fmt.Println(len(c), cap(c))

	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	// time.Sleep(1 * time.Second)

	select {
	case n := <-ch:
		fmt.Println(n)
	case <-time.After(16 * time.Microsecond):
		fmt.Println("After")
		// default:
		// 	fmt.Println("Ничего")
	}

	fmt.Println("Finished!")
}
