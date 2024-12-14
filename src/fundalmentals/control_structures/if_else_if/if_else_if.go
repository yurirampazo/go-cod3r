package main

import "fmt"

func main() {
	fmt.Println(9.2)
	fmt.Println(8.2)
	fmt.Println(7.2)
	fmt.Println(6.2)
	fmt.Println(5.2)
	fmt.Println(4.2)
}

func grades(n float32) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n <= 9 {
		return "B"
	} else if n >= 7 && n <= 8 {
    return "C"
	} else if n >= 6 && n <= 7 {
		return "D"
	} else if n >= 5 && n <= 6 {
		return "E"
	}
	return "F"

}
