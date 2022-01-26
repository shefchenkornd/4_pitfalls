package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const pathJsonFile = "52_marshal_struct_to_json/output.json"

type Professor struct {
	Name       string     `json:"name"`
	ScientId   int        `json:"scient_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Bob",
		ScientId:  1,
		IsWorking: true,
		University: University{
			Name: "MTSU",
			City: "Rostov-on-Don",
		},
	}

	// 1. Превратим профессора в последовательность байт с отступами
	byteArr, err := json.MarshalIndent(prof1, "", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal. Error: %v", err)
	}

	fmt.Println(string(byteArr)) // {"name":"Bob","scient_id":1,"is_working":true,"university":{"name":"MTSI","city":"Rostov-on-Don"}}

	err = os.WriteFile(pathJsonFile, byteArr, 0664) // -rw-rw-r
	if err != nil {
		log.Fatalf("Failed to create file. Error: %v", err)
	}
	fmt.Println("JSON file successfully created!")
}
