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


func main() {
	fmt.Printf("Average: %.2f\n", average(6.7, 8.9, 0.5, 0.5, 0.3, 10.0))
	fmt.Printf("Average: %.2f\n", average())
}