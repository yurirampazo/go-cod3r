package main

import "fmt"


func main() {

	//Arrays are static.
	//Arrays are only support with same data types

	var grades [3] float64
	fmt.Println(grades)


	grades[0], grades [1], grades[2] = 7.8, 4.3, 9.1
	// grades[3] = 1
	fmt.Println(grades)

	total := 0.0 
	for i := 0; i < len(grades); i++ {
		total += grades[i]
 }

 averageGrade := total / float64(len(grades))
 fmt.Printf("Average %.2f\n", averageGrade)

}