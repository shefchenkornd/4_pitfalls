package main

import (
	"fmt"
	"sync"
)

// Приведение типов
func main() {
	var i int = 10
	var a int32 = 10
	fmt.Printf("%v - %T \n", i, i) // 10 - int
	fmt.Printf("%v - %T \n", a, a) // 10 - int32

	var c int32
	c = int32(i)

	fmt.Println(c == a) // true

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
