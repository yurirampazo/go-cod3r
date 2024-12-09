package main

import "fmt"


func main() {
	fmt.Print("Same")
	fmt.Print(" line")


	fmt.Println(" New")
	fmt.Println("line")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("X value is:" + xs)
	fmt.Println("X value is:" , xs)

	fmt.Printf("X float value if %.2f\n", x)
	fmt.Printf("X String value os %s\n", xs)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d \n%f \n%.1f", a, b, b)
	fmt.Printf("\n%03d", a)
	fmt.Printf("\n%s", d)
	fmt.Printf("\n\n%t, \n%b, \n%v, \n %s", c,c,c,c)
	fmt.Println()

}