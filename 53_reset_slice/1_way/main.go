package main

import "fmt"

const sizeCap = 20

// Setting the slice to `nil`, однако этот метод изменяет емкость до нуля.
func main() {
	sl := make([]int, 0, sizeCap)
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
	fmt.Printf("Slice: %v \nLength: %d \nCapacity: %d \nPoint: %p \n", sl, len(sl), cap(sl), sl) // len:3, cap:20, point: 0xc0000b8000
	fmt.Println()

	sl = nil
	fmt.Printf("Slice: %v \nLength: %d \nCapacity: %d \nPoint: %p \n", sl, len(sl), cap(sl), sl) // len:0, cap:0, point: 0xc0000b8000
	fmt.Println()

	// Обратите внимание, что во втором случае (sl = nil) ёмкость слайса стала равной 0.
}
