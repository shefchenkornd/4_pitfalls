package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(file.Name())

	b := make([]byte, 100)
	var count int
	count, err = file.Read(b)
	if err != nil {
		if io.EOF != err {
			log.Fatal(err)
		}
	}

	fmt.Printf(
		"read %d bytes: %q\n",
		count,
		b[:count],
	)

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Home directory for particular user:", dir)
}
