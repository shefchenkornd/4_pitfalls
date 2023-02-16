package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)


type message struct {
	cond *sync.Cond
	msg  string
}

// Пример
// Предположим у нас есть некоторый общий ресурс в системе. Одна группа процессов
// может изменять его состояния, а другая группа должна реагировать на эти изменения.
func main() {
	msg := message{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// 1
	for i := 1; i <= 3; i++ {
		go func(num int) {
			for {
				msg.cond.L.Lock()
				msg.cond.Wait()
				fmt.Printf("hello, i am worker%d. text:%s\n", num, msg.msg)
				msg.cond.L.Unlock()
			}
		}(i)
	}

	// 2
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		msg.cond.L.Lock()
		msg.msg = scanner.Text()
		msg.cond.L.Unlock()

		msg.cond.Broadcast()
	}

}
// Ответ:
// Мы запустили 3 goroutine которые ждут сигнала. Обратите внимание, что вызов Wait не просто блокирует, он приостанавливает текущую процедуру, позволяя другим процедурам запускаться.
// При входе Wait вызывается Unlock в Locker переменной Cond, а при выходе из Wait вызывается Lock в Locker переменной Cond. К этому нужно немного привыкнуть.
// Во второй части мы читаем ввод из консоли и отправляем сигнал об изменении состояния.
// Broadcast отправляет сигнал всем ожидающим goroutine. А метод Signal находит goroutine, которая ждала дольше всего и будет ее.