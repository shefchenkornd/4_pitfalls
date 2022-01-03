package main

import "fmt"

// Функцию recover() можно использовать для поимки/перехвата panic. Это получится, только если вызвать её в блоке defer.
func main() {
	// Примечание: recover() должна быть выше по коду, чем panic(). Потому что исполнение кода дальше panic() не происходит!
	defer func() {
		fmt.Println("recovered:",recover()) // recovered: not good
	}()

	panic("not good")
}
