package main

import "fmt"

var i32 int32 = 0
var i64 int64 = 0

func main() {
	// fmt.Println(i32 == i64) // error: Invalid operation: i32 == i64 (mismatched types int32 and int64)

	fmt.Println(eq(i32, i64))
}


func eq(val1 interface{}, val2 interface{}) bool {
	return val1 == val2
}