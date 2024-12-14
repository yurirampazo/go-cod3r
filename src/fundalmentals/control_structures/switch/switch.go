package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(9.8)
	fmt.Println(7.8)
	fmt.Println(4.8)

	withoutParameter()
	fmt.Println(types(32))
	fmt.Println(types("32"))
	fmt.Println(types(func() {}))
	fmt.Println(types(true))
}

// Most common
func grades(n float64) string {
	var grade = int(n)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade"
	}
}

// Without parameter
func withoutParameter() {
	t := time.Now()

	switch { //Finds the first true, more like an if
	case t.Hour() > 5 && t.Hour() < 12:
		fmt.Println("Good Morning!")
		fmt.Println("Good Morning!")
	case t.Hour() > 11 && t.Hour() < 18:
		{
			fmt.Println("Testing more than one statement")
			fmt.Println("Good Afternoon!")
		}
	default:
		fmt.Println("Good night!")
		fmt.Println("Good night!")
	}
}

// Advanced
func types(i interface{}) string {

	switch i.(type) {
	case int:
		return "INTEGER"
	case float32, float64:
		return "FLOAT"
	case string:
		return "STRING"
	case func():
		return "FUNCTION"
	default:
		return "CALL TypeOf FUNCTION from Reflect import"
	}
}
