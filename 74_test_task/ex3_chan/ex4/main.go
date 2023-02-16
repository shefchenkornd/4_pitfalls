package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0..9
	}

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			_, ok := alreadyStored[doubles[i]]
			if !ok {
				alreadyStored[doubles[i]] = struct{}{} //


				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	
	close(uniqueIDs)
	for val := range uniqueIDs {
		fmt.Println(val)
	}
}
