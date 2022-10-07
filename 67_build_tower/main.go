package main

import (
	"fmt"
	"strings"
)

func main() {
	res := TowerBuilder(3)
	fmt.Println("Result = ", res)
}

// TowerBuilder is build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors.
// A tower block is represented with "*" character.
// [
//  "     *     ",
//  "    ***    ",
//  "   *****   ",
//  "  *******  ",
//  " ********* ",
//  "***********"
// ]
func TowerBuilder(nFloors int) []string {
	totalStars := (nFloors * 2) - 1

	var res []string
	star := 1
	for i := 1; nFloors >= i; i++ {
		floor := strings.Repeat("*", star)
		spaceCount := (totalStars - star) / 2
		if spaceCount > 0 {
			floor = strings.Repeat("_", spaceCount) + floor + strings.Repeat("_", spaceCount)
		}

		res = append(res, floor)
		fmt.Println(floor)
		star +=2
	}

	return res
}
