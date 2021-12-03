package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int = 10
	var a int32 = 10
	fmt.Printf("%v - %T \n", i, i)
	fmt.Printf("%v - %T \n", a, a)

	var c int32
	c = int32(i)

	fmt.Println(c == a)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
