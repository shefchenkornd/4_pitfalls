package main

import (
	"fmt"
	"runtime"
	"time"
)

// Утечка горутины в заблокированном канале при отсутствии default в select{}
func main() {
	fmt.Println("Goroutine на старте:", runtime.NumGoroutine()) // 1 горутина
	ch := make(chan struct{})


	go func() {
		select {
		case <-ch:
			fmt.Println("получаем из канала")
		// default:							// если отключить default, то у нас эта анонимная горутина утечёт!!!
		// 	fmt.Println("default")
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Goroutine на финише:", runtime.NumGoroutine()) // 2 горутина
}
