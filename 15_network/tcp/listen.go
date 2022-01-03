package main

import (
	"fmt"
	"log"
	"net"
)

// Для прослушивания и приемы входящих запросов в пакете net определена функция net.Listen:
func main() {
	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
