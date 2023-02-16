package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const URL = "https://google.com"

// Структура http.Client имеет ряд полей, которые позволяют настроить ее поведение:
// Timeout: устанавливает таймаут для запроса
// Jar: устанавливает куки, отправляемые в запросе
// Transport: определяет механиз выполнения запроса
func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}
