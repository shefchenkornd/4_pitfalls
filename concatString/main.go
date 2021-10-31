package main

import (
	"bytes"
	"fmt"
	"strings"
)

var params = []string{"string-1", "string-2", "string-3", "string-4"}

func main() {
	fmt.Println(cPlus(params))
	fmt.Println(cBytes(params))
	fmt.Println(cStrings(params))
	fmt.Println(cCopy(params))
}

func cCopy(args []string) string {
	lenB := make([]byte, len(args)*len(args[0]))
	offset := 0
	for _, str := range args {
		fmt.Println(offset)
		offset += copy(lenB[offset:], str)
	}

	return string(lenB[:])
}

func cPlus(args []string) string {
	var result string

	for _, str := range args {
		result += str
	}

	return result
}

func cBytes(args []string) string {
	buffer := bytes.Buffer{}

	for _, str := range args {
		buffer.WriteString(str)
	}

	return buffer.String()
}

func cStrings(args []string) string {
	builder := strings.Builder{}

	for _, str := range args {
		builder.WriteString(str)
	}

	return builder.String()
}
