package main

import (
	"fmt"
	"log"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, 1024 * 4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		source := string(input[0:n])
		// на основании полученных данных получаем из словаря перевод
		target, ok := dict[source]
		if ok == false {
			target = "undefined"
		}

		// выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)

		// отправляем данные клиенту
		conn.Write([]byte(target))
	}
}

func main() {
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
		go handleConnection(conn)
	}
}
