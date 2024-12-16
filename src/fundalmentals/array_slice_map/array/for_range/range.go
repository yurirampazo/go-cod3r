package main

import "fmt"

func main() {
	numbers := [...]int{1,2,3,4,5}  //compiler infers the numbers of indexes

	for i, item := range numbers {
			fmt.Printf("%d) %d\n" , i+1, item)
	}

	fmt.Println("Ignoring index")

	for _, item := range numbers {
		fmt.Printf("%d\n", item)
}






}