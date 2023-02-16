package main

import (
	"fmt"
	"sync"
)

// Go. Задачи по concurrency. Часть 1
func main() {
	writes := 1000
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mx := sync.RWMutex{} // Используем RWMutex{}, потому что у нас есть и запись, и чтение!

	// ЗАПИСЬ
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mx.Lock()
			defer mx.Unlock()
			storage[i] = i
		}()
	}

	// ЧТЕНИЕ
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mx.RLock()
			defer mx.RUnlock()
			_, _ = storage[i]
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
