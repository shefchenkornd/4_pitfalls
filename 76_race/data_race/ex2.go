package main

import (
	"errors"
	"fmt"
	"sync"
)

var amountErr = errors.New("amount is more than accFrom.balance")

type account struct {
	balance int
}

func transfer(mx *sync.Mutex ,from, to *account, amount int) error {
	if amount > from.balance {
		return amountErr
	}

	mx.Lock()
	from.balance -= amount
	to.balance += amount
	mx.Unlock()

	return nil
}

var number = 10000

func main() {
	from, to := &account{balance: number}, &account{balance: 0}

	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for i := 0; i < number; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			err := transfer(&mx, from, to, 1)
			if err != nil {
				fmt.Printf("error for n=%d\n", n)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("From.balance", from.balance)
	fmt.Println("To.balance", to.balance)
}
