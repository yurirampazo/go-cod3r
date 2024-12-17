package main

import "fmt"

func main() {

	s1 := make([]int , 10, 20)
	s2 := append(s1,1, 2, 3,)

	fmt.Println(s1, s2)

	s1[0] = 7  // points to both internal array
	fmt.Println(s1, s2)



	//append and copy: onlyh  applied to slices

	array1 := [3]int{1, 2, 3}
	var slice []int

	// array1 = append(array1, 4, 5, 6)
	slice = append(slice, 4, 5, 6)

	fmt.Println(array1, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)
	fmt.Println(slice2)


}