package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"

	conn, err := net.Dial("tcp", "go.dev:80")
	if err != nil {
		log.Fatalln("Error connect to web-site:", err)
	}

	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		log.Fatalln("Error while write:", err)
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
