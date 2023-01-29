package main

import (
	"fmt"
	"sync"
)

var counter int = 0 // Общий ресурс

/**
На уровне кода мьютекс представляет тип sync.Mutex.
Для блокирования доступа к общему разделяемому ресурсу у Mutex вызывается метод Lock(),
а для разблокировки доступа - метод Unlock().
 */
func main() {
	ch := make(chan bool) // канал

	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for i := 1; i < 5; i++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock() // деблокируем доступ
	ch <- true
}
