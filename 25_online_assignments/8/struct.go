package main

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}
}

// Fill in the blanks for lines marked with A and B, to ensure that the printed output is "foo"
func main() {
	p:= f()
	print(p.m) // print "foo"
}
