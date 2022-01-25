package main

import "fmt"

func main() {
	// UTF-8 and string literals
	//
	// That means that when we store a character value in a string, we store its byte-at-a-time representation
	// 	Here’s a simple program that prints a string constant with a single character three different ways,
	// once as a plain string, once as an ASCII-only quoted string, and once as individual bytes in hexadecimal.
	// To avoid any confusion, we create a “raw string”, enclosed by back quotes, so it can contain only literal text.
	// (Regular strings, enclosed by double quotes, can contain escape sequences as we showed above.)
	const placeOfInterest = `⌘` // the “Place of Interest” symbol ⌘

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest) // ⌘
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest) // "\u2318"
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ { // e2 8c 98
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	// Source code in Go is defined to be UTF-8 text; no other representation is allowed
	// In short, Go source code is UTF-8, so the source code for the string literal is UTF-8 text.
	// To summarize, strings can contain arbitrary bytes, but when constructed from string literals, those bytes are (almost always) UTF-8.
}
