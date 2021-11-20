package main

import (
	"fmt"
)

// Метод Write предназначен для копирования данных их среза байт p в определенный ресурс - файл, сетевой интерфейс и т.д.
// Метод возвращает количество записанных байтов и объект ошибки.

type phoneWriter struct{}

func (ph phoneWriter) Write(bs []byte) (int, error) {
	countNumber := 0
	for _, r := range bs {
		if r >= 48 && r <= 59 {
			fmt.Print(string(r))
			countNumber++
		}
	}

	return countNumber, nil
}

func main() {
	bytes := []byte("+1(234)567 9010")

	ph := phoneWriter{}
	ph.Write(bytes)
}
