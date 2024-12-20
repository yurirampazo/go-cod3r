package main

import "fmt"

type grade float64

func (g grade) between(begin, end float64) bool {
	return float64(g) >= begin && float64(g) <= end
}

func gradeForConcept(g grade) string {

	if g > 10 {
		g = 10.0
	}
	
	switch {
	case g.between(9.0, 10.0):
		return "A"

	case g.between(7.0, 8.99):
		return "B"
	
	case g.between(5.0, 6.99):
		return "C"

	case g.between(3.0, 4.99):
		return "D"
	
	case g.between(1.0, 2.99):
		return "E"

	default:
		return "F"

	}
}

func main() {
	fmt.Println(gradeForConcept(0.2))
	fmt.Println(gradeForConcept(1.2))
	fmt.Println(gradeForConcept(2.2))
	fmt.Println(gradeForConcept(3.2))
	fmt.Println(gradeForConcept(4.2))
	fmt.Println(gradeForConcept(5.2))
	fmt.Println(gradeForConcept(6.2))
	fmt.Println(gradeForConcept(7.2))
	fmt.Println(gradeForConcept(8.2))
	fmt.Println(gradeForConcept(9.2))
}