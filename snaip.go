package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := [][]int{
		[]int{},
	}

	res := Snail(sl)
	fmt.Println("Result: ", res)
}

// Snail
// [1,2,3,6,9,8,7,4,5]
func Snail(snaipMap [][]int) []int {
	var result []int

	for {
		keys := make(map[int]int, len(snaipMap))
		for i, item := range snaipMap {
			// right
			if i == 0 {
				for _, v := range item {
					result = append(result, v)
				}
				keys[i] = 0
				continue
			}

			lastIndx := len(snaipMap) - 1
			if i != lastIndx {
				// down
				internalLastIndx := len(snaipMap[i]) - 1
				result = append(result, snaipMap[i][internalLastIndx])

				keys[i] = internalLastIndx
			} else {
				// to the left
				for y := len(snaipMap[i]) - 1; y >= 0; y-- {
					result = append(result, snaipMap[i][y])
				}

				keys[i] = 0
			}
		}

		if len(keys) > 0 {
			// bubble by keys indexes
			var slKeys []int
			for i := range keys {
				slKeys = append(slKeys, i)
			}
			sort.Ints(slKeys)

			for i := len(slKeys) - 1; i >= 0; i-- {
				key, _ := keys[i]

				if i == 0 && key == 0 {
					// delete snaipMap[0]
					copy(snaipMap[0:], snaipMap[1:])
					snaipMap = snaipMap[:len(snaipMap)-1]
				} else if key == 0 {
					// delete last element in snaipMap
					snaipMap = snaipMap[:len(snaipMap)-1]
				} else if i != 0 && key != 0 {
					// delete last element in subSlice
					snaipMap[i] = snaipMap[i][:key]
				}
			}
		}

		if len(snaipMap) == 0 {
			break
		}
	}

	return result
}
