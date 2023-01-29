package main

import (
	"math/rand"
	"sync"
)

// fatal error: concurrent map read and map write
// В таком случае нужно использовать что-то из списка ниже:
//  * sync.Mutex
//  * sync.RWMutex
//  * sync.Map
func main() {
	data := make(map[int]int)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				// read
				if data[rand.Int()] != 0 { // fatal error: concurrent map read and map write
					//
				}
			}
		}(&wg)
	}

	// write
	for i := 0; i < 100; i++ {
		data[rand.Int()] = rand.Int()
	}
}
