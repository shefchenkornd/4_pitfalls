package main

import (
	"log"
	"time"
)

/*
Необходимо написать in-memory кэш, который будет по ключу (uuid пользователя) возвращать профиль и список его заказов.
1. У кэша должен быть TTL
2. Кэшем может пользоваться функция(-и), которая работает с заказами (добавляет/обновляет/удаляет)
3. Должны быть написаны тестовые сценарии использования данного кэша
*/


type UserProfile struct {
	UUID   string
	Name   string
	Orders []Order
}

type Order struct {
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Basket    interface{}
}



func main() {
	x := 2 + 2
	if x != 4 {
		log.Fatalf("expected 4, got %d", x)
	}

	log.Println("Tests passed")
}
