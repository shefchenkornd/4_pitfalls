package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// slice - это структура
/*
type slice struct {
	  array unsafe.Pointer &[10]int[]
	  len int 0
	  cap int 10
}
*/
func main() {
	s1 := "abc"
	s2 := "def"
	fmt.Println(s1 + s2)

	sl := []int{1, 2, 3}
	fmt.Println(sl)

	// m := make(map[rune]int)
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Println(stats.HeapAlloc)

	bigArray := [1 << 25]int{}
	_ = bigArray

	runtime.ReadMemStats(stats)
	fmt.Println(stats.HeapAlloc)

	sl3 := []int{1, 22, 3}

	sl2 := sl[1:]
	sl[1] = 55
	sl2[0] = 22
	fmt.Println(sl2)
	fmt.Println(sl)

	fmt.Println(reflect.DeepEqual(sl, sl3))

	// nums := make([]int, 1, 2) // len:1, cap:2
	// fmt.Println(len(nums), cap(nums)) // [0]
	// fmt.Println(nums) // [0]
	//
	// appendSlice(nums, 1024)
	// fmt.Println(nums) // [0]
	//
	// mutateSlice(nums, 1, 512)
	// fmt.Println(nums) // panic: index out of range
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
