package main

import "fmt"


// Consider the following code:
func main() {
 var y = 458
 var p = &y

 *p = 500
 fmt.Println(y)
}

// Answer: 500