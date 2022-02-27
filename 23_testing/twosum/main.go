package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Printf("Result: %#v\n", res)

	return nil
}

func twoSum(nums []int, target int) []int {
	for index, number := range nums {
		for index2, number2 := range nums {
			if number+number2 == target {
				return []int{index, index2}
			}
		}
	}

	return []int{}
}
