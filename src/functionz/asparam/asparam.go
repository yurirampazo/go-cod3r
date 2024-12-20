package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func exec(fun func(int, int) int, p1, p2 int) int {
	return fun(p1, p2)
}

func main() {
result	:= exec(multiply, 3, 4)
fmt.Println(result)
}