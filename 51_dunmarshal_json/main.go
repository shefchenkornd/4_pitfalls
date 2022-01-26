package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"log"
	"os"
)

const pathToJsonFile = "51_dunmarshal_json/users.json"

// Decode (unmarshal) JSON to map
func main() {
	// 1. Создаём файл дескриптор
	file, err := os.Open(pathToJsonFile)
	if err != nil {
		log.Fatalf("Failed to open file %v. Error: %v", pathToJsonFile, err)
	}
	defer file.Close()
	fmt.Println("File descriptor successfully created!")

	// 2. Вычитываю набор байт из файл дескриптора
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read error: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal, error: %v", err)
	}
	fmt.Println("Unmarshal is successfully!")

	fmt.Println("Result:")
	spew.Dump(result)
}
