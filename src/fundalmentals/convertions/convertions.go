package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	//be carefull.. this is catchin value from ASCII table.
	fmt.Println("Test " + string(97))


	//int to string
	fmt.Println("Test " + strconv.Itoa(123))


	//string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	//string to boolean
	b,_ := strconv.ParseBool("true")
	a,_ := strconv.ParseBool("1")
	if b && a {
		fmt.Println(b)
	}

}