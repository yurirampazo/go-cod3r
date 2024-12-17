package main

import "fmt"

func main() {

	// maps must be initialized
	// var approved map[int]string
	approved := make(map[int]string)

	approved[12345678] = "maria"
	approved[21345678] = "ana"
	approved[55345678] = "pedro"


	fmt.Println(approved)

	for doc, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, doc)
	}

	fmt.Println(approved[12345678])
	delete(approved, 12345678)
	fmt.Println(approved[12345678])


	// ---

	employeesAndSalaries := map[string]float64{
		"yuri": 10.13,
		"gaby": 11.13,
		"marico": 13.13,
	}

	employeesAndSalaries["yuri"] = 15.02
	delete(employeesAndSalaries, "no_one")

	for name, salary := range employeesAndSalaries {
		fmt.Println(name, salary)
	}
		



}