package main

import "fmt"

func inc1(n int) {
	n++
}

// revision: pointer indicated by *
func inc2(n *int) {
	// revision2: *int here is used to dereference.
	// Access value of the pointer
	*n++
}

func main() {
	n := 1
	
	inc1(n)
	fmt.Println(n)

	/** Revision: & gets the address of variable, 
	using with n in this case, will pass the variable 
	value reference, not only the value.
	It'll be incremented
	*/
	inc2(&n)
	fmt.Println(n)


}
