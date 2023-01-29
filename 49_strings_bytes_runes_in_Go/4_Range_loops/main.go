package main

import (
	"fmt"
	"unicode/utf8"
)

// for range loop
func main() {
	// Есть только один способ, которым Go обрабатывает UTF-8 специально, а именно при использовании цикла for для строки
	fmt.Println("Печатаем \"abc\" через цикл for{}")
	for i, i2 := range "abc" {
		fmt.Println(i, string(i2))
	}

	fmt.Println("Печатаем \"абв\" через цикл for{}")
	for i, i2 := range "абв" {
		fmt.Println(i, string(i2))
	}
	// P.s: обратите внимание, что индексы в обоих циклах выше идут по разному!


	const nihongo = "日本語"

	// A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration.
	// Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes, and the code point is its value.
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// The output shows how each code point occupies multiple bytes:
	// U+65E5 '日' starts at byte position 0
	// U+672C '本' starts at byte position 3
	// U+8A9E '語' starts at byte position 6

	// -OR-

	// Here is a program equivalent to the for range example above,
	// but using the DecodeRuneInString function from `unicode/utf8` package to do the work.
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	// U+65E5 '日' starts at byte position 0
	// U+672C '本' starts at byte position 3
	// U+8A9E '語' starts at byte position 6
}
