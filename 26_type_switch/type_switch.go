package main

import (
	"fmt"
	"log"
	"strconv"
)

type Option struct {
	params []string
}

// Как проверить тип переменной во время выполнения на языке Go?
// У Go есть специальная форма переключателя, предназначенная для проверки тип переменной (она называется переключателем типа)
func main() {
	var op = &Option{}
	var i int64 = 64

	params := []interface{}{"abc", 123, i}

	for _, param := range params {
		fmt.Println(param)

		err := setOptions(op, param)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println(op)
}

func setOptions(option *Option, param interface{}) error {
	switch v := param.(type) {
	default:
		return fmt.Errorf("unexpected type %T \n", v)
	case string:
		option.params = append(option.params, v)
	case int:
		option.params = append(option.params, strconv.Itoa(v))
	}

	return nil
}
