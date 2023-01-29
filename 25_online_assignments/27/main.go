package main

import "fmt"

type S struct{}

func (s S) f() {}

func InitPointer() *S {
	var s *S
	return s
}

func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

// Consider the following code:
func main() {
	fmt.Println("InitPointer() == nil | Result:", InitPointer() == nil)
	fmt.Println("InitEfaceType() == nil | Result:", InitEfaceType() == nil)
	fmt.Println("InitEfacePointer() == nil | Result:", InitEfacePointer() == nil)

	fmt.Println()
	fmt.Printf("%T, %v \n", InitPointer(), InitPointer())
	fmt.Printf("%T, %v \n", InitEfaceType(), InitEfaceType())
	fmt.Printf("%T, %v \n", InitEfacePointer(), InitEfacePointer())
}

// Answer:
// InitPointer() == nil | Result: true
// InitEfaceType() == nil | Result: false
// InitEfacePointer() == nil | Result: false