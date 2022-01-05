package main

import (
	"runtime"
	"time"
)

var _ = runtime.GOMAXPROCS(3)

var a, b int

func u1() {
	a = 1
	b = 2
}

func u2() {
	a = 3
	b = 4
}

func p() {
	println(a)
	println(b)
}

// Go может менять порядок некоторых операций, но общее поведение внутри горутины, где это происходит, не меняется.
// Однако сказанное не относится к порядку исполнения самих горутин.
// Если запустить этот код несколько раз, то можно увидеть различные комбинации переменных a и b
//
// Если нужно сохранить порядок операций чтения и записи среди нескольких горутин, то используйте каналы или соответствующие конструкции из пакета sync.
func main() {
	go u1()
	go u2()
	go p()
	time.Sleep(1 * time.Second)
}
