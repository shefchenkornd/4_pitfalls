package main

import (
	"fmt"
	"os"
)

// Отложенные вызовы исполняются в конце содержащей их функции, а не в конце содержащего их кодового блока.
func main() {
	var targets = []string{"path1", "path2", "path3"}
	for _, target := range targets {
		func() { // Один из способов решения проблемы — обернуть кодовый блок в функцию.
			f, err := os.Open(target)
			if err != nil {
				fmt.Println("bad target:", target, "error:", err)
				return
			}
			defer f.Close() // ok
			// сделай что-нибудь с файлом...
		}()
	}
}
