package main

import (
	"fmt"
	"math"
	"strconv"
)

// 89 = 8^1 + 9^2
// 135 = 1^1 + 3^2 + 5^3
func main() {
	a, b := uint64(1), uint64(200)
	res := SumDigPow(a, b)
	fmt.Println("Result = ", res)
}

func SumDigPow(a, b uint64) []uint64 {
	// это решение из каты
	r := []uint64{}
	for i := a; i <= b; i++ {
		s := 0
		p := strconv.Itoa(int(i))
		for l, n := range p {
			z, _ := strconv.Atoi(string(n))
			s += int(math.Pow(float64(z), float64((l + 1))))
		}
		if uint64(s) == uint64(i) {
			r = append(r, uint64(s))
		}
		s = 0
	}
	return r


	// это моё решение
	//var eureka []uint64
	//var temp uint64
	//
	//for i := a; b >= i; i++ {
	//	temp = 0
	//
	//	for indx, r := range strconv.FormatUint(i, 10) {
	//		s := math.Pow(float64(r-'0'), float64(indx+1))
	//		temp += uint64(s)
	//	}
	//
	//	if i == temp {
	//		eureka = append(eureka, i)
	//	}
	//}
	//
	//return eureka
}
