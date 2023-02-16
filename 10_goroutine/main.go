package main

import (
	"time"
)

/*
Свойства канала:
  1. потокобезопасна (hchan mutex)
  2. хранение элементов, FIFO (hchan buf
  3. передача данных в горутин (sendDirect, operations with buf)
  4. блокировка горутин (sendq/recvq, sudog | calls to scheduler: gopark(), goready() )
  5. каналы создаются в Heap
  6. под "капотом" канал использует Mutex
  7. при чтении из канала мы получаем КОПИЮ, т.е. у sender'a будет своя копия, а у reader своя копия
  8. если мы читаем из reader, он блокируется, и потом приходит sender, то доступ к стеку другой горутины

  ----------------------------------------------------
   Из чего состоит канал:
   type hchan struct {
   	qcount	uint  — кол-во эл-тов, которое хранится в буфере
   	dataqsiz uint — размерность буфера

   	buf *buffer — ссылка на сам буфер
   	closed uint32 — флаг, который говорит о том закрыт ли канал? (здесь используется sync/atomic, поэтому uint32)

   	elemsize uint16 —
   	elemtype uint32 — тип элемента

   	recvq *list — указатели на связанный список из горутин, которые ожидают чтения или записи
   	sendq *list — аналогично пункту выше

   	recvx uint — receive index | номер ячейки буфера, из которых будет производиться на чтение/запись данных
   	sendx uint — send index | номер ячейки буфера, из которых будет производиться на чтение/запись данных

   	lock mutex — защищает все поля в hchan
   }

   ПО СУТИ КАНАЛ ЭТО УКАЗАТЕЛЬ НА HCHAN СТРУКТУРУ!
*/
func main() {
	ch := make(chan uint32, 4)

	// включи режим "debug" и посмотри как изменяется `ch`
	ch <- 1
	_ = <-ch

	ch <- 2
	ch <- 3

	go func() {
		for {
			ch <- 45
		}
	}()

	time.Sleep(1 * time.Second)

	return

}
