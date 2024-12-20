package main

import "fmt"

func closure() func() {
	x := 10
	var fun = func ()  {
		fmt.Println(x)
	}
	return fun
}

func main() {
	x := 20
	fmt.Println(x)

	printX := closure()
	printX()
}