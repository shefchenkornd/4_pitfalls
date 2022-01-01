package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Когда вы декодируете/десериализуете JSON-данные в интерфейс,
// Go по умолчанию обращается с числовыми значениями в JSON как с числами float64. Значит, вот такой код вызовет panic:
func main() {
	var data = []byte(`{"status":200}`)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result["status"]) // 200
	fmt.Printf("%T \n", result["status"]) // float64

	// Неправильно:
	var status = result["status"].(int) // ОШИБКА: panic: interface conversion: interface {} is float64, not int
	fmt.Println("status value:",status)

	// Правильно:
	status2 := uint64(result["status"].(float64))
	fmt.Println(status2) // 200
	fmt.Printf("%T \n", status2) // uint64
}
