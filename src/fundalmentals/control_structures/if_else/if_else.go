package main

import "fmt"

func main() {
	printResult(7.3)
	printResult(5.3)
}

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade", grade)
	} else {
		fmt.Println("Reproved with grade", grade)
	}
}
