package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Буферизированный ввод-вывод
// Чтение через буфер
func main() {
	file, err := os.Open("some.dat")
	if err != nil {
		log.Fatalln("Unable to open file:", err)
	}
	defer file.Close()

	b := bufio.NewReader(file)
	for {
		str, err := b.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		}

		fmt.Print(str)
	}
}
