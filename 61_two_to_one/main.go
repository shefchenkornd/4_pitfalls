package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	res := TwoToOne(a, b)

	fmt.Println(res)
}

// TwoToOne
// Take 2 strings s1 and s2 including only letters from `a` to `z`.
// Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.
// a = "xyaabbbccccdefww"
// b = "xxxxyyyyabklmopq"
// longest(a, b) -> "abcdefklmopqwxy"
func TwoToOne(s1 string, s2 string) string {
	var builder strings.Builder

	for _, str := range []string{s1, s2} {
		for _, l := range str {
			if strings.Contains(builder.String(), string(l)) == false {
				builder.WriteString(string(l))
			}
		}
	}

	res := strings.Split(builder.String(), "")
	sort.Strings(res)

	return strings.Join(res, "")
}
