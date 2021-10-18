package main

import "fmt"

type DBStruct struct {
	DBConnectionName string
}

func (ms *DBStruct) PrintDBConnectionName() {
	fmt.Println(ms.DBConnectionName)
}

func getDBStruct() (*DBStruct, error) {
	return nil, fmt.Errorf("oh noes, couldn't create the DB connection")
}

func main() {
	/**
	Здесь пропустили ошибку и значение `s` пустое. Но оказалось, что код не будет паниковать до строчки #10

	Потому что Go спокойно вызывает функции в нулевых структурах.
	Но в этом опасность пойти по ложному пути при отладке. Функция вызывается нормально,
	так что со структурой все в порядке и проблема в чем-то другом.

	Лучшее решение — проявить бдительность и обязательно проверить наличие нулевых указателей.
	*/
	s, _ := getDBStruct()
	s.PrintDBConnectionName()
}
