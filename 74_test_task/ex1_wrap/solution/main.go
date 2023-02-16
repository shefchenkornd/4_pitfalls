package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	// fmt.Println(rnd)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc() (int64, error) {
	ch := make(chan int64)

	go func(ch chan int64){
		ch <- unpredictableFunc()
	}(ch)

	select{
	case <-time.After(10*time.Second):
		return 0, fmt.Errorf("Some error")
	case res := <-ch:
		return res, nil
	}
}

func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}