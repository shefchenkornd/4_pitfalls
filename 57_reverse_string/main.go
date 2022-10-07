package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "double  spaced  words"
	res := ReverseWords(str)

	fmt.Println(res)
}

// ReverseWords
// Complete the function that accepts a string parameter, and reverses each word in the string.
// All spaces in the string should be retained.
// Example: "This is an example!" ==> "sihT si na !elpmaxe"
func ReverseWords(str string) string {
	sl := strings.Split(str, " ")

	var builder strings.Builder
	for _, word := range sl {
		builder.WriteString(doReverseWords(word))
		builder.WriteString(" ")
	}

	return strings.TrimSuffix(builder.String(), " ")
}

func doReverseWords(str string) string {
	runes := []rune(str)
	var builder strings.Builder

	for i := len(runes) - 1; 0 <= i; i-- {
		builder.WriteRune(runes[i])
	}

	return builder.String()
}
