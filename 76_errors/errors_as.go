package main

import (
	"errors"
	"fmt"
)

type FileNotFound struct {
	err error
}

func (r *FileNotFound) Error() string {
	return r.err.Error()
}

func (r *FileNotFound) Unwrap() error {
	return r.err
}

func todo() *FileNotFound {
	return &FileNotFound{err: errors.New("file not found")}
}

func main() {
	err := todo()
	var e *FileNotFound
	fmt.Println(errors.As(err, &e)) // проверяем, имеет ли ошибка определенный type!
}
