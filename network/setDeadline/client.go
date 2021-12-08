package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			log.Fatalln("Некорректный ввод", err)
		}

		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			log.Fatalln(err)
		}

		// получаем ответ
		fmt.Print("Перевод:")
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				log.Fatalln(err)
				break
			}

			fmt.Println(buff[0:n])
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
		fmt.Println()
	}
}
