package main

import "fmt"

func main() {
	results := make(map[int]int)
	ch := make(chan struct{})

	go factorials(5, ch, results)
	<-ch

	for i, v := range results{
		fmt.Println(i, " - ", v)
	}
}


func factorials(n int, ch chan struct{}, results map[int]int){
	defer close(ch)     // закрываем канал после завершения горутины
	result := 1
	for i:=1; i <= n; i++{
		result *= i
		results[i] = result
	}
}