package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i = struct {}{}		// 0 bytes
	// var i bool				// 1 bytes
	// var i [1]int8 			// 1 bytes

	// var i [1]int16 			// 2 bytes

	// var i rune				// 4 bytes // alias for int32
	// var i [1]int32 			// 4 bytes

	// var i float64 = 32		// 8 bytes
	// var i map[string]int		// 8 bytes
	// var i func()				// 8 bytes
	// var i [1]int 			// 8 bytes
	// var i [1]int64 			// 8 bytes

	// var i string				// 16 bytes
	// var i complex128 		// 16 bytes
	// var i interface{} 		// 16 bytes

	// var i []int				// 24 bytes
	// var i []string			// 24 bytes

	fmt.Println(
		unsafe.Sizeof(i),
	)
}
