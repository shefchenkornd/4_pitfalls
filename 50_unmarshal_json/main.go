package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const pathToJsonFile = "50_unmarshal_json/users.json"

// Users struct for representation total slice
// First level of Json object parsing
type Users struct {
	Users []User
}

// User Internal user representation
// Second level of Json object parsing
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

func PrintUser(u *User) {
	fmt.Println()
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social: VK: %s and FB: %s \n", u.Social.Vkontakte, u.Social.Facebook)
}

// Decode (unmarshal) JSON to struct
func main() {
	file, err := os.Open(pathToJsonFile)
	if err != nil {
		log.Fatalf("Failed to open file %v. Error: %v", pathToJsonFile, err)
	}
	defer file.Close()
	fmt.Println("File descriptor successfully created!")

	var users Users

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read error: %v", err)
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("Failed to unmarshal, error: %v", err)
	}
	fmt.Println("Unmarshal is successfully!")

	for _, user := range users.Users {
		PrintUser(&user)
	}
}
