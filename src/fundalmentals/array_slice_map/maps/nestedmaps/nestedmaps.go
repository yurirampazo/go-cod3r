package main

import "fmt"

func main() {
	peopleByInitial := map[string]map[string]float64{
		"A": {
			"ana": 15.04,
			"agnaldo": 13.04,
		},
		"J": {
			"Jana": 10.98,
		},
		"P": {
			"paulo": 200.98,
		},
	}

	for initial, person := range peopleByInitial {
		fmt.Println(initial, person)
	}

	fmt.Println("---")
	delete(peopleByInitial["A"], "ana")
	for initial, person := range peopleByInitial {
		fmt.Println(initial, person)
	}

	fmt.Println("---")
	delete(peopleByInitial, "P")
	for initial, person := range peopleByInitial {
		fmt.Println(initial, person)
	}


}