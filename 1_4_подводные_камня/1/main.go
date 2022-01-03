package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

/**
Это странная особенность циклов на Go. Когда вы ссылаетесь на переменную цикла, сама она ссылается на значения массива.
При этом фактически ссылка идет на последнее значение, на которое эта переменная цикла ссылалась.
Простое решение — придать циклу следующий вид:
*/
func getUserIDs(users []User) []*int {
	ids := make([]*int, 0, len(users))
	for _, user := range users {
		id := user.Id // <- Здесь он принимает копию значения, значения переменной цикла, и ссылается на него.
		ids = append(ids, &id)
	}

	return ids
}

func main() {
	users := []User{
		{1, "Andy"}, {2, "John"}, {3, "Jane"},
	}

	ids := getUserIDs(users)
	for _, id := range ids {
		fmt.Println(*id)
	}

}
