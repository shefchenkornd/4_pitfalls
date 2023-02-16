package main

import (
	"fmt"
)

// что выведет данная функция?
func main() {
	msgChan := make(chan string, 1)
	close(msgChan)
	stopChan := make(chan struct{})
	close(stopChan)

	select {
	case msgChan <- "msg": // если попадём в этот case, то будет panic
		fmt.Println("msg sent")
	case <-stopChan: // если попадём в этот case, то всё ОК
		fmt.Println("stop signal received")
	}
}