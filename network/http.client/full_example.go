package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Добавляем заголовки
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-agent", "MSIE/15.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
