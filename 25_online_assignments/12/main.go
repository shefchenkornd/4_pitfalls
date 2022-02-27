package main

func f() (int, int) {
	return 1, 2
}

// Assume `x` is declared and `y` is not declared. Which of the following clauses are correct (choose one or more)?
func main() {
	var x int = 5
	var y int

	x, y = f() // или так, или как указано ниже
	//x, _ = f()

	print(x, y)
}
