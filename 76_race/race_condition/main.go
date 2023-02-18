package main

import (
	"fmt"
	"time"
)

// Race condition — это недостаток, возникающий, когда время или порядок событий влияют на правильность программы.
//
// Важно, что Race condition — это семантическая ошибка.
//
// Помочь могут хорошие практики и проверенные паттерны.
func main() {
	// case#1
	// ex1()

	// case#2
	ex2()

}

// Результат выполнения кода зависит от порядка выполнения горутин. Это типичная ошибка race condition
// Мы ожидаем:
// A->B
// A->B
//
// Но в реальности будет что-то подобное:
// A->B
// B
// A->A->A->B
// A->B
func ex1() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("A->")
		}()

		go func() {
			fmt.Printf("B \n")
		}()
	}

	time.Sleep(2 * time.Second)
}

// В результате на консоле получим четные и нечетные числа, а расчитывали увидеть только четные.
// Проблемы с доступом к общим ресурсам проще обнаружить автоматически и решаются они обычно с помощью синхронизации
func ex2() {
	timeout := time.After(100 * time.Nanosecond)
	x := 0

	for {
		go func() {
			x++
		}()

		go func() {
			if x%2 == 0 {
				time.Sleep(10 * time.Nanosecond)
				fmt.Println(x)
			}
		}()

		select {
		case <-timeout:
			return
		default:
		}
	}
}
