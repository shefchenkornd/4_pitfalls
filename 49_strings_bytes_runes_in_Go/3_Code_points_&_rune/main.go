package main

import (
	"fmt"
	"unicode/utf8"
)

// “code point”
//
// The Unicode standard uses the term “code point” to refer to the item represented by a single value.
// The code point U+2318, with hexadecimal value 2318, represents the symbol ⌘. (For lots more information about that code point, see its Unicode page.)
//
// The Unicode code point U+0061 is the lower case Latin letter ‘A’: a
// letter à - code point (U+00E0)
// In general, a character may be represented by a number of different sequences of code points, and therefore different sequences of UTF-8 bytes.
// The concept of character in computing is therefore ambiguous, or at least confusing, so we use it with care. To make things dependable, there are normalization techniques.

// rune
//
// “Code point” is a bit of a mouthful, so Go introduces a shorter term for the concept: rune.
// The term appears in the libraries and source code, and means exactly the same as “code point”, with one interesting addition.
func main() {
	// напечатает символ
	fmt.Println("⌘") // ⌘


	// напечатает руну
	fmt.Println('⌘') // 8984


	// 	строчная латинская буква "a" в rune('a') имеет значение 97 (int32 base)
	//	теперь переводим это из int32 в hex base и получаем 61
	//	что соответствует кодовой точке Unicode U+0061.
	fmt.Println(rune('a')) // 97
	fmt.Printf("%x \n", rune('a')) // 61 или кодовая точка Unicode U+0061

}
// To summarize, here are the salient points:
//
// Go source code is always UTF-8.
// A string holds arbitrary bytes.
// A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
// Those sequences represent Unicode code points, called runes.
// No guarantee is made in Go that characters in strings are normalized
