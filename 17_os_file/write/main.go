package main

import (
	"log"
	"os"
)

func main() {
	text := "This new line for file"
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal("Unable to create file:", err)
	}
	defer file.Close()

	// Для записи текстовой информации в файл можно применять метод WriteString() объекта os.File, который заносит в файл строку:
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Write to file:", err)
	}


	bFile, err := os.Create("hello.bin")
	if err != nil {
		log.Fatal("Unable to create binary file:", err)
	}
	defer bFile.Close()

	// Для записи нетекстовой бинарной информации в виде набора байт применяется метод Write() (реализация интерфейса io.Writer):
	data := []byte("binary data")
	println(data)
	_, err = bFile.Write(data)
	if err != nil {
		log.Fatal("Write to binary file:", err)
	}
}
