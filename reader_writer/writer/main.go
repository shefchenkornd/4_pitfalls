package main

import (
	"fmt"
	"strings"
)

// Метод Write предназначен для копирования данных их среза байт p в определенный ресурс - файл, сетевой интерфейс и т.д.
// Метод возвращает количество записанных байтов и объект ошибки.

type phoneWriter struct{
	phone string
}

func (ph phoneWriter) Write(bs []byte) (int, error) {
	countNumber := 0
	var sb strings.Builder

	for _, r := range bs {
		if r >= 48 && r <= 59 {
			sb.WriteString(string(r))

			fmt.Print(string(r))
			countNumber++
		}
	}

	ph.phone = sb.String()

	return countNumber, nil
}

func main() {
	bytes := []byte("+1(234)567 9010")

	ph := phoneWriter{}
	ph.Write(bytes)

	fmt.Println(ph.phone)
}
