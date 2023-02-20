package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Livelock – ситуация, при которой два или более потока не могут выполнять полезной работы по причине борьбы за общий ресурс.
Это похоже на deadlock, но разница в том, что процессы становятся “вежливыми” и позволяют другим делать свою работу (поэтому live, а не dead).

Простые примеры:
  1. Двое встречаются лицом к лицу. Каждый из них пытается посторониться, но они не расходятся, а несколько секунд сдвигаются в одну и ту же сторону.
  2. Вы делаете телефонный звонок, но человек на другом конце тоже пытается вам позвонить. Вы оба повесите трубку и попробуйте снова через одно и то же время, что снова создаст такую же ситуацию. Это может продолжаться вечность.

Обнаружить Livelock труднее, чем deadlock, просто потому, что может показаться, что программа работает. Она может реагировать на сигналы,
потреблять ресурсы и как то менять состояния, но выйти из цикла и завершить работу уже не в состоянии.
*/


type spoon struct {
	owner *diner
}

func (s spoon) use() {
	fmt.Printf("%s has eaten!\n", s.owner.name)
}

type diner struct {
	name     string
	isHungry bool
}

func (d *diner) eatWith(sp *spoon, spouse *diner) {
	for d.isHungry {

		// 1
		if sp.owner != d {
			time.Sleep(1 * time.Second)
			continue
		}

		// 2
		if spouse.isHungry {
			fmt.Printf("%s: You eat first my darling %s!\n", d.name, spouse.name)
			sp.owner = spouse
			continue
		}

		// 3
		sp.use()
		d.isHungry = false
		fmt.Printf("%s: I'm stuffed? my darling %s!\n", d.name, spouse.name)
		sp.owner = spouse
	}
}

/*
Рассмотрим простой пример livelock:
муж и жена пытаются поужинать, но между ними только одна ложка.
Каждый из супругов слишком вежлив, и передает ложку, если другой еще не ел.

Обедаем пока не утолим голод(isHungry=false).
  * Если ложка сейчас не у нас, то подождем
  * Если супруг(а) голодна, то уступим и передадим ложку ему/ей
  * Используем ложку и наконец-то обедаем

Поесть этим милым людям не суждено. До третьего блока выполнение не дойдет.
*/
func main() {
	husband := &diner{
		name:     "Bob",
		isHungry: true,
	}
	wifi := &diner{
		name:     "Alice",
		isHungry: true,
	}

	sp := &spoon{owner: husband}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		husband.eatWith(sp, wifi)
		wg.Done()
	}()

	go func() {
		wifi.eatWith(sp, husband)
		wg.Done()
	}()

	wg.Wait()
}
