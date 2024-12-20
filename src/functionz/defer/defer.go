package main

import "fmt"

func getApprovedValue(num int) int {
	defer fmt.Println("end!")
	if num > 5000 {
		fmt.Println("High value")
		return 5000
	}
	fmt.Println("Low value")
	return num

}

func main() {
	fmt.Println(getApprovedValue(6000))
}
