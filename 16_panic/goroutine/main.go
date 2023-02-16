package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Пример:
// если есть горутина в которой случилась panic(), то эта паника выйдет за пределы горутины и ВСЁ приложение завершится с ошибкой
// чтобы этого избежать нужно добавить в горутину следующий код, который обработает панику в области defer:
// defer func() {
// 	if err := recover(); err != nil {
// 		log.Println(err)
// 	}
// }()
// Совет: defer recovery() нужно добавлять в каждую горутину, чтобы предотвратить падение сервиса.
func main() {
	createUser()
}

func createUser() {
	fmt.Println("Начало!")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() { // РЕШЕНИЕ ПРОБЛЕМЫ!!!!!!
			if err := recover(); err != nil {
				log.Println("Перехватили панику:", err)
			}
		}()
		err := sendEmail()
		fmt.Println(err)
	}()

	wg.Wait()

	time.Sleep(3 * time.Second)
	fmt.Println("Конец!")

}

func sendEmail() error {
	panic("a problem")

	return fmt.Errorf("some error: %v", 555)
}
