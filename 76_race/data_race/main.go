package main

import (
	"fmt"
	"sync"
)

type account struct {
	balance uint64
}

func deposit(acc *account, amount uint64) {
	// wrong!
	acc.balance += amount

	// Right. Иногда более эффективным решением будет использовать пакет atomic.
	// atomic.AddUint64(&acc.balance, amount)
}

// Data race — это состояние когда разные потоки обращаются к одной ячейке памяти без какой-либо синхронизации и как минимум один из потоков осуществляет запись.
func main() {
	wg := sync.WaitGroup{}

	acc := &account{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			deposit(acc, 1)
		}()
	}

	wg.Wait()

	// Ожидаем баланс равный 1000руб, но получим около 984руб, т.е. мы потеряли часть денег!
	// потому что мы не использовали Mutex или пакет Atomic.
	fmt.Println(acc.balance)
}
// Причина в том, что операция acc.balance += amount не атомарная. Она может разложиться на 3 шага:
// tmp := acc.balance
// tmp = tmp + amount
// acc.balance = tmp
//
// Пока мы меняем временную переменную в одном потоке, в других уже изменен основной balance.
// Таким образом теряется часть изменений.
// У Go есть хороший Data Race Detector с помощью которого такие ошибки можно обнаружить [https://go.dev/doc/articles/race_detector]
