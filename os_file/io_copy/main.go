package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wr, err := io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal("io.Copy", err)
	}
	fmt.Println(wr, err)
}
