package main


import "fmt"

func average(numbers ...float64) float64 {
	if len(numbers) == 0  || numbers == nil {
		return 0
	}


	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func printApproved(approved ...string) {
	fmt.Println("Approved list")

	for i, person := range approved {
		fmt.Printf("%d) %s\n", i+1, person)
	}
}


func main() {
	fmt.Printf("Average: %.2f\n", average(6.7, 8.9, 0.5, 0.5, 0.3, 10.0))
	fmt.Printf("Average: %.2f\n", average())

	fmt.Println("---")

	approved := make([]string, 0)
	approved = append(approved, "Goku", "Gohan", "Shin")
	printApproved(approved...)


}