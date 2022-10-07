package main

import (
	"fmt"
	// "regexp"
	// "sort"
	// "strconv"
	"strings"
)

func main() {
	a := "is2 Thi1s T4est 3a"
	res := Order(a)

	fmt.Println(res)
}

// Order
// Your task is to sort a given string. Each word in the string will contain a single number.
// This number is the position the word should have in the result.
//
// Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).
//
// If the input string is empty, return an empty string. The words in the input String will only contain valid consecutive numbers.
// Example: "is2 Thi1s T4est 3a"  -->  "Thi1s is2 3a T4est"
func Order(sentence string) string {
	a := strings.Split(sentence, " ")
	r := make([]string, len(a))
	for _, st := range a {
		for _, c := range st {
			if c >= '0' && c <= '9' {
				r[int(c-'1')] = st
				break
			}
		}
	}

	return strings.Join(r, " ")

	// if sentence == "" {
	// 	return ""
	// }
	//
	// pieces := strings.Split(sentence, " ")
	//
	// re := regexp.MustCompile("[0-9]")
	//
	// m := make(map[int]string)
	// for _, word := range pieces {
	// 	number := re.FindString(word)
	// 	if number != "" {
	// 		n, _ := strconv.Atoi(number)
	// 		m[n] = word
	// 	}
	// }
	//
	// keys := make([]int, 0, len(m))
	// for key := range m {
	// 	keys = append(keys, key)
	// }
	//
	// sort.Ints(keys)
	//
	// var result string
	// for _, key := range keys {
	// 	result += m[key] + " "
	// }
	//
	// return strings.TrimSuffix(result, " ")
}
