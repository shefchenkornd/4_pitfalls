package db

import "fmt"

var (
	NameDb, Connect = "Postgress", 0
)

func HelloDatabase() {
	fmt.Println("I'm Database function")
}
