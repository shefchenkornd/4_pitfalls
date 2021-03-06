package main

import (
	"fmt"
	"unicode/utf8"
)

// Неизменяемость строк
//
// Если вы попытаетесь обновить отдельные символы строковой переменной с помощью оператора индекса, то это не сработает.
// Строки — это байт-слайсы (byte slices), доступные только для чтения. Если вам все-таки нужно обновить строку,
// то стоит использовать байт-слайс и преобразовывать его в строку по необходимости.
func main() {
	a := "abcd"

	s := []byte(a)
	s[0] = 65 // <-код большой буквы "A"

	fmt.Println(string(s)) // Abcd

	// вот с русскими буквами не знаю как правильно работать
	b := "абв"

	u := []byte(b)

	for i, r := range b {
		fmt.Println(i, r, b[i])
	}

	u[0] = 65 // <-код большой буквы "A"

	// todo: узнать как работать с кир
	fmt.Println(string(u)) // A�бв

	// Строки — не всегда текст в кодировке UTF-8
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) // true

	data2 := "A\xBC"
	fmt.Println(utf8.ValidString(data2)) // false

	// Длина строк
	data3 := "♥"
	fmt.Println("len:", len(data3)) // выводит: 3 | Встроенная функция len() возвращает не символ, а количество байт
	fmt.Println("Rune:", utf8.RuneCountInString(data3)) // выводит: 1 | Технически функция RuneCountInString() не возвращает количество символов, потому что один символ может занимать несколько рун

	data4 := "é"
	fmt.Println(len(data4))                    // выводит: 3
	fmt.Println(utf8.RuneCountInString(data4)) // выводит: 2
}
