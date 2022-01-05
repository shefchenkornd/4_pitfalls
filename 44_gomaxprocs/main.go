package main

import (
	"fmt"
	"runtime"
)

// The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously.
func main() {
	fmt.Println(runtime.GOMAXPROCS(-1)) // выводит: 8 (1 on play.golang.org)
	fmt.Println(runtime.NumCPU())       // выводит: 8 (1 on play.golang.org)
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1)) // выводит: 20
}
