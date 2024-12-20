package main

import "fmt"

// Anonymous function
var sum = func(a, b int) int {
	return a + b
}


func main() {
	fmt.Println(sum(2, 3))

	sub := func(a, b int) int {
		return a -b
	}

	fmt.Println(sub(2,3))
}