package main

import "fmt"


func main() {
	ch := make(chan int, 1000)

	ch <- 1
	ch <- 2
	ch <- 3


	// Планировщик видит, что у нас нет др. горутин, только горутина main() и понимает, что нам никто
	// никогда не отправит закрытие канал и поэтому говорит Deadlock

	// close(ch) // чтобы этого избежать закрой канал!!!
	for val := range ch {
		fmt.Println(val)
	}
}
