package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting HTTP Server by localhost:8081...")
	http.HandleFunc("/", GetIndex)
	http.HandleFunc("/header", GetHeader)

	fmt.Println("HTTP Server started successfully")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

// GetIndex return info about index page
// w - responseWriter (куда писать ответ)
// r - request (откуда брать запрос)
func GetIndex(w http.ResponseWriter, req *http.Request)  {
	byteArray := []byte("It is simple HTTP Server. Current datetime: " + time.Now().Format("02-Jan-2006 15:04:05"))
	w.Write(byteArray)
	// -or- то есть писать данные можно, как в сам `w http.ResponseWriter`, так и через функцию fmt.Fprintf()
	// fmt.Fprintf(w, "It is simple HTTP Server")
}

func GetHeader(w http.ResponseWriter, req *http.Request)  {
	for headerName, value := range req.Header {
		for _, h := range value {
			fmt.Fprintf(w, "%v: %v\n", headerName, h)
		}
	}
}