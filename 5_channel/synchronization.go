package main

import "fmt"

func main() {
	results := make(map[int]int)
	sync := make(chan struct{})

	go factorials(5, sync, results)
	<-sync

	for i, v := range results{
		fmt.Println(i, " - ", v)
	}
}


func factorials(n int, sync chan struct{}, results map[int]int){
	defer close(sync)     // закрываем канал после завершения горутины
	result := 1
	for i:=1; i <= n; i++{
		result *= i
		results[i] = result
	}
}