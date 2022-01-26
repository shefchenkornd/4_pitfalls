package main

import "fmt"

const sizeCap = 20

// Используем полное выражение слайса (full slice expression) с указанием ёмкости исходного слайса
func main() {
	sl := make([]int, 0, sizeCap)
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
	fmt.Printf("Slice: %v \nLength: %d \nCapacity: %d \nPoint: %p \n", sl, len(sl), cap(sl), sl) // len:3, cap:20
	fmt.Println()

	sl = sl[:0:cap(sl)]
	fmt.Printf("Slice: %v \nLength: %d \nCapacity: %d \nPoint: %p \n", sl, len(sl), cap(sl), sl) // len:0, cap:20
	fmt.Println()

	// Обратите внимание, ёмкость исходного слайса стала прежней.
}
