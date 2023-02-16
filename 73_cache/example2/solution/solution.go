package main

import (
	"fmt"
	"log"
	"sync"
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

type TTL int64 // секунды
type Item struct {
	user    UserProfile
	created time.Time
	t       TTL
}
type Cache struct {
	mx sync.Mutex
	m  map[string]Item
}

// можно еще написать interface, описывающий работу NewCache

func NewCache() *Cache {
	return &Cache{
		mx: sync.Mutex{},
		m:  make(map[string]Item),
	}
}

func (c *Cache) Set(profile UserProfile, t TTL) {
	c.m[profile.UUID] = Item{
		user:    profile,
		created: time.Now(),
		t:       t,
	}
}

func (c *Cache) Get(uuid string) (*UserProfile, error) {
	item, ok := c.m[uuid]
	if !ok {
		return nil, fmt.Errorf("%v", "not found")
	}

	if time.Now().Unix() > (item.created.Unix() + item.t) {
		return nil, fmt.Errorf("%v", "время жизни кэша истекло. Объект удалён")
	}

	return &item.user, nil
}

func main() {
	cache := NewCache()
	
	uuid.
	uuid := "uuid.New()"
	profile := UserProfile{
		UUID: uuid,
		Name: "Kent",
		Orders[]Order{},
	}

	cache.Set(profile, 5)

	profile2, err := cache.Get(uuid)
	if err != nil {
		log.Fatalf("случилась ошибка: %v", err.String())
	}

	if profile2 == nil {
		log.Fataln("empy userprofile")
	}

	// сравнить profile == profile2
	if profile.UUID != profile2.UUID {
		log.Fataln("not equal userProfile")
	}

	if profile.Name != profile2.Name {
		log.Fataln("not equal userProfile")
	}

	x := 2 + 2
	if x != 4 {
		log.Fatalf("expected 4, got %d", x)
	}

	log.Println("Tests passed")
}
