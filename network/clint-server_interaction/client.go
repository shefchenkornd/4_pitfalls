package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":4545")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for {
		var source string
		fmt.Println("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}

		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			log.Fatalln(err)
		}

		// получаем ответ
		fmt.Print("Перевод:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}

		fmt.Println(string(buff[:n]))
	}
}
