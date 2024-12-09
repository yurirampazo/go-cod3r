package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// int numbers
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal int is", reflect.TypeOf(3200))
	

	// only positive - u from unsigned
	/*
	uint8 -> byte
	uint16 -> short
	utin32 -> int
	uint64 -> long
	*/

	var test string = "test"
	fmt.Println(test)
	
	var b byte = 12
	fmt.Println(b)

	var s int16 = 12
	var ss uint16 = 12
	fmt.Println(s, ss)

	i1 := math.MaxInt64
	fmt.Println(i1)
	fmt.Println(reflect.TypeOf(i1))
	
	
	var i2 rune = 'a' // represents an unicode mapping od the value in format int32
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.09
	fmt.Println("Type is", reflect.TypeOf(x))
	fmt.Print("Default value for float number is", reflect.TypeOf(42.90)) //float64


	//boolean
	bo := true //does not accept integer
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(bo)


	// string
	s1 := "Hello, my name is Yuri"
	fmt.Println(s1 + "!")
	fmt.Println("The size of string is", len(s1))
	
	// multi line strings
	s2 := `
	Hello
	My
	Name
	Is
	Yuri
	`
	fmt.Println(s2)
	fmt.Println("The size of string is", len(s2))


	//char??
	char := 'a'
	fmt.Println(reflect.TypeOf(char)) //int32


}