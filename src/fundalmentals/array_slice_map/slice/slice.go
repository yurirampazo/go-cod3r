package main

import (
	"fmt"
	 "reflect"
)

func main() {
		// array
		numbers := [...]int{1, 2, 3, 4, 5} //compiler infers the numbers of indexes
		// numbers = numbers.append(1) cannot do this

		// slice
		numbers2 := []int{1, 2, 3}
		numbers2 = append(numbers2, 4, 5)
	
		for i, item := range numbers {
			fmt.Printf("%d) %d\n", i+1, item)
		}
	
		fmt.Println("Ignoring index")
	
		for _, item := range numbers2 {
			fmt.Printf("%d\n", item)
		}
	
		fmt.Println(reflect.TypeOf(numbers))
		fmt.Println(reflect.TypeOf(numbers2))
}