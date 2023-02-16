package main

import "fmt"

// Number
// Объявим тип интерфейса Number для использования в качестве ограничения типа.
// Объявим объединение int64 и float64 внутри интерфейса
type Number interface {
	int64 | float64
}


func sumInt64(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloat64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// используем Number тип ограничение в generic функции
func sumGenerics[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"1": 10,
		"2": 20,
	}

	floats := map[string]float64{
		"1": 35.9,
		"2": 26.8,
	}

	fmt.Println("Non-Generic Sums =", sumInt64(ints), sumFloat64(floats))

	// all right
	fmt.Println("Generic Sums:", sumGenerics[string, int64](ints), sumGenerics[string, float64](floats))

	// можно опустить типы аргументов в вызове функции
	// Обратите внимание, что это не всегда возможно. Например, если вам нужно вызвать универсальную функцию без аргументов, вам нужно будет включить аргументы типа в вызов функции.
	fmt.Println("Generic Sums, type parameters inferred:", sumGenerics(ints), sumGenerics(floats))


	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
