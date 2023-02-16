package main

type Swimmer interface {
	Swim(string)
}

type Walker interface {
	Walk(int) int
}

// Animal
// Встраиваем Walker и Swimmer интерфейсы в Animal интерфейс
type Animal interface {
	Walker
	Swimmer
}


type Duck struct {
	Animal
}

func (*Duck) Swim(path string)  {
	
}

func (*Duck) Walk(distance int) int {
	return 0
}

func main() {
	var duck Animal = &Duck{}
	duck.Swim("")
	duck.Walk(100)
}
