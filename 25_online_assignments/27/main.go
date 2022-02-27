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
	fmt.Println(InitEfacePointer())


	print(InitPointer() == nil, " ")
	print(InitEfaceType() == nil, " ")
	print(InitEfacePointer() == nil)
}

// Answer: true false false