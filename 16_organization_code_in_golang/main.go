package main

import (
	"./db"
	logger "./log"
	"fmt"
)

// Организация кода в Go
func main() {
	db.HelloDb()
	db.HelloDatabase()
	fmt.Println(db.NameDb)

	logger.HelloLog()
}
