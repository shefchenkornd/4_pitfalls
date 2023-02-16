package main

import (
	"fmt"
	"log"
	"net/http"
)

const URL = "https://ya.ru"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status, "\n")
	fmt.Println("Header:", resp.Header, "\n")

	fmt.Println("Body:", resp.Header)
	for {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		if n == 0 || err != nil {
			break
		}

		fmt.Println(string(bs[:n]))
	}
}
