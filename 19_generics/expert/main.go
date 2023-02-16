package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(max(1, 5))
	fmt.Println(max(12.21, 79.00))
}

/*
вместо написания собственных ограничений,
можно воспользоваться стандартным пакетом "golang.org/x/exp/constraints"
от самих разработчиков.
*/
func max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}
