package main

import (
	"fmt"
)

func main() {

	i := 1

	for i <= 10 { // GO does not support while, this is the while here
		fmt.Println(i)
		i++
	}


	for i := 0; i <= 10; i += 2 {
		fmt.Printf("%d", i)
	}

	fmt.Println("\nMixing")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	//infinity loop
	// for {
	// 	fmt.Println("4 Ever")
	// 	time.Sleep(time.Second)
	// }

	//ForEach in the arrays section


}