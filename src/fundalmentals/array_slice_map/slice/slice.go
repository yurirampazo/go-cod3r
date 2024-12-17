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


		a1 := [3]int{1, 2, 3} // array
		s1 := []int{1, 2, 3} // slice

		fmt.Println(a1, s1)
		fmt.Println(reflect.TypeOf(a1))
		fmt.Println(reflect.TypeOf(s1))

		a2 := [5]int{1, 2, 3, 4, 5}
		s2 := a2[1:3]

		fmt.Println(a2, s2)
		s3 := a2[:2]

		fmt.Println(s3)

		s4 := s2[0:1]
		fmt.Println(s4)



}