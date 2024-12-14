package main

import "fmt"

//Like kotlin, there is no need for ternary in GO, we can do it with if else
func getResult(grade float64) string {
	// return grade >= 6 ? "Cooking" : "Cooked";

	if grade >6{
		return "Approved"
	} 
	return "reproved"
}

func main() {
	fmt.Println(getResult(6.2))
}