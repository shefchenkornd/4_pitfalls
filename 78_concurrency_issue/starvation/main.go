package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Starvation — это любая ситуация, когда параллельный процесс не может получить все ресурсы, необходимые для выполнения его работы.
При livelock все параллельные процессы одинаково “голодают”, и никакая работа не выполняется до конца.

Пример:
У нас будет два работника. Один жадный(greedyWorker), другой вежливый(politeWorker). Обоим дается одинаковое кол-во времени на их полезную работу — спать по 3 наносекунде.
greedyWorker жадно удерживает общий ресурс(sharedLock) на протяжении всего цикла работы, тогда как politeWorker пытается блокировать его только тогда, когда это необходимо.

За одно и то же время, жадный работник получил почти вдвое больше возможностей выполнять свою работу и владеть общим ресурсом.
*/
func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()

	wg.Wait()
}
