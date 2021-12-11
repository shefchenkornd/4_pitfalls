package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const N = 10

// Which of the following is the correct way to fix this issue so that the program avoids the "concurrent map writes" error?
func main() {
	m := make(map[int]int)
	var mx sync.Mutex // answer: add

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(mx *sync.Mutex) {
			defer mx.Unlock()  // answer: add
			defer wg.Done()

			mx.Lock()  // answer: add
			m[rand.Int()] = rand.Int()
		}(&mx)
	}

	wg.Wait()
	println(len(m))
	fmt.Println("map:", m)
}
// Answer: create mutex lock for m[rand.Int()] = rand.Int()