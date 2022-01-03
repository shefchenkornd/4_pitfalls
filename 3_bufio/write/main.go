package main

import (
	"bufio"
	"log"
	"os"
)

// Буферизированный ввод-вывод
// Запись через буфер
func main() {
	rows := []string{
		"Hello Go!",
		"Welcome to Go!",
	}

	file, err := os.Create("some.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	b := bufio.NewWriter(file)
	for _, str := range rows {
		b.WriteString(str)
		b.WriteString("\n")
	}

	b.Flush()
}
