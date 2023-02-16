package main

import "fmt"

// каков будет результат? и почему?
func main() {
	sl := []string(nil)


	sl = append(sl , []string(nil)...)
	fmt.Println(sl)
	fmt.Println(len(sl), cap(sl)) // len = 0, cap = 0

	sl = append(sl , "")
	fmt.Println(sl)
	fmt.Println(len(sl), cap(sl))
}
