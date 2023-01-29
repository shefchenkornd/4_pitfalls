package main

import (
	"fmt"
	"unsafe"
)

// unexpected fault address 0x498bea
// fatal error: fault
// [signal SIGSEGV: segmentation violation code=0x2 addr=0x498bea pc=0x480ef4]
func main() {
	// a := "Andrey Berenda" // будет error, который выше
	a := string([]byte("Andrey Berenda")) // код отработает успешно!
	fmt.Println(a)

	b := toSlice(a)
	b[5] = 'i'
	fmt.Println(a)
}

func toSlice(a string) []byte {
	return *(*[]byte)(unsafe.Pointer(&a))
}
