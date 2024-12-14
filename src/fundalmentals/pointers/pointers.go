package main

import "fmt"

func main() {
	i := 1

	//GO does not support arithmetic operations on pointers

	var p *int = nil  //declaring the p pointer of some future variable
	
	p = &i //retrieving the variable i address and setting in the pointer

	*p++ //*p dereference the pointer and gets the value where p is pointing
	 //then we can increase or decrease

	 fmt.Println(p, *p, i, &i)


	 fmt.Println(*p == i) //true 2 2
	 fmt.Println(p == &i) //true 0xc0000ae010 0xc0000ae010

	 i++
	 fmt.Println(i, *p)


}