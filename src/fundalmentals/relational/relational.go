package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Test" == "Test")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1 == d2)

	type Person struct {
		Name string
	}

	p1 := Person{"Yuri"}
	p2 := Person{"Yuri"}

	//THese are two referenced memory variables, but we are comparing the content
	fmt.Println("Are equals?", p1 == p2)

}