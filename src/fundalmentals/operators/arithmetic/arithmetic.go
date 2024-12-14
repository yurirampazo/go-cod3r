package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a + b)
	fmt.Println("Diference =>", a - b)
	fmt.Println("Multiplication =>", a * b)
	fmt.Println("Division =>", a / b)
	fmt.Println("Division module =>", a % b)

	//bitwise
	fmt.Println("E =>", a&b) // 11 & 10 = 10
	fmt.Println("Ou =>", a|b) // 11 | 10 = 11
	fmt.Println("Xor ->", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//math
	fmt.Println("Greater =>", math.Max(float64(a), float64(b)))
	fmt.Println("Shorter =>", math.Min(c, d))
	fmt.Println("Power =>", math.Pow(c, d))

}