package main

// Consider the following code:
// Which of the following modifications to the code will allow the program to exit the outer for-loop?
func main() {
	outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break outer
		}
		println()
	}
	// Answer: put `outer` at 6 line and `break outer` at line 10
}
