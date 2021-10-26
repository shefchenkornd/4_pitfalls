package main

import "fmt"

/**
с помощью этой конструкции мы проверяем имплементирует ли структура repo интерфейс Repo
Если не имплементирует, то получим вот такую ошибку:
cannot use (*repo)(nil) (type *repo) as type Repo in assignment:
	*repo does not implement Repo (missing UpdateUser method)
 */
var _ Repo = (*repo)(nil)

type User struct {
	ID string
}

// Repo is a repository interface
type Repo interface {
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
}

// repo is an implementation of above Repo interface
type repo struct {}

// GetUser gets user from DB by id
func (r *repo) GetUser(id string) (*User, error) {
	return nil, nil
}

// UpdateUser updates user in DB by id
func (r *repo) UpdateUser(user *User) error {
	return nil
}

func main() {
	fmt.Println("Что это за конструкция???")
}
