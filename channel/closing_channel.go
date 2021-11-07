package main

import "fmt"

func main() {
	intCh := make(chan int, 3)

	intCh <- 3
	intCh <- 7
	close(intCh)

	/**
	Первое значение - val - это собственно данные из канала, а opened представляет логическое значение, которое равно true, если канал открыт и мы можем успешно считать из него данные.
	Например, мы можем проверять с помощью условной конструкции состояние канала:
	 */
	for i := 0; i < cap(intCh); i++ {
		v, opened := <-intCh
		if opened {
			fmt.Println(v)
		} else {
			fmt.Println("Канал закрыт!")
		}
	}
}
