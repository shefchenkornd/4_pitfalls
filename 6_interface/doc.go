package main

/*
Рекомендации:
  1. минималистичность интерфейса
  2. независимость от реализации
  3. интерфейсы лучше размещать в месте использования
  4. с помощью интерфейсов можем уменьшить связность между модулями/пакетами
  5. с случае добавления нового хранилища (Storage, например, в виде какого-то cache), нам не нужно реализовать все методы
  6. Простота тестирования (с помощью //go:generate go run github.com/vektra/mockery/v2@v2.18.0 --name=UserProvider
*/

type User struct {
	ID   int
	Name string
	Age  int
}

/*
Wrong!
потому что, этот Storage слишком ОБЩИЙ!!!
В случае добавления новой реализации в виде cache, придётся реализовать все методы, а нам этого не нужно!
*/

type Storage interface {
	users() ([]User, error)
	UsersByAge(age int) ([]User, error)
	User(id int) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int) error
}

/*
Right
Интерфейсы, которые представлены ниже, удовлетворяют рекомендация, которые описаны выше!
*/

//go:generate go run github.com/vektra/mockery/v2@v2.18.0 --name=UserProvider
type UserProvider interface {
	User(id int) (User, error)
}

type UserCreator interface {
	Create(user User) error
}
