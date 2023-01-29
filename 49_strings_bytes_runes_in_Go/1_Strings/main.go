package main

import "fmt"

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

// In Go, a string is in effect a read-only slice of bytes.
// As we saw, indexing a string yields its bytes, not its characters: a string is just a bunch of bytes.
//
// In fact, the definition of “character” is ambiguous, and it would be a mistake to try to resolve the ambiguity by defining that strings are made of characters
func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample) // ��=� ⌘

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ { // bd b2 3d bc 20 e2 8c 98
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Print with %x:")
	fmt.Printf("%x\n", sample) // bdb23dbc20e28c98

	// A nice trick is to use the “space” flag in that format, putting a space between the % and the x.
	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample) // bd b2 3d bc 20 e2 8c 98

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample) // "\xbd\xb2=\xbc ⌘"

	// If we are unfamiliar or confused by strange values in the string, we can use the “plus” flag to the %q verb.
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample) // "\xbd\xb2=\xbc \u2318"
}
























