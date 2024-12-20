package main

import "fmt"

func change(p1, p2 int) (second, first int) {
	second = p2
	first = p1
	return // clean return
}

func main() {
	x, y := change(2,3)
	fmt.Println(x, y)
}

