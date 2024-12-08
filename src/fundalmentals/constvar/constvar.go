package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1215
	var radius = 3.2 // type float 64 infered 

	//shorter declaration
	area := PI * m.Pow(radius, 2)
	fmt.Println("Area is =", area)
	
	const (
		a = 1
		b= 2
	)
	
	var (
		c = 3
		d = 4
	)
	
	fmt.Println(a, b, c, d,)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := false, "testing", 10
	fmt.Println(g, h, i)





}