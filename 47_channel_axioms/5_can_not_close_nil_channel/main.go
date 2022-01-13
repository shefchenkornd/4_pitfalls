package main

func main() {
	var c chan bool

    close(c) // panic: close of nil channel
}
