package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type resp struct {
	id  int
	err error
}

/*
context можно закрыть:
  * cancel()
  * timeout
  * deadline
*/
func main() {
	rand.Seed(time.Now().Unix())
	ch := make(chan resp)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go RPCCall(ctx, ch)

	resp := <-ch
	fmt.Printf("ID - %d  | error - %v \n", resp.id, resp.err)

	cancel()
}

func RPCCall(ctx context.Context, ch chan resp) {
	select {
	case <-ctx.Done(): // в данном случае сработает context.Timeout, те. истекут наши 2*time.Second
		ch <- resp{
			id:  0,
			err: errors.New("timeout"),
		}
	case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
		ch <- resp{
			id: rand.Int(),
		}
	}
}
