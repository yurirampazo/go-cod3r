package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d", n)
	case n < 2:
		return 1, nil
	default:
		nextFactorial, _ := factorial(n - 1)
		return n * nextFactorial, nil	
	}
}

func simpleFactorial(n uint) uint {
	switch {
	case n < 2:
		return 1
	default:
		return n * simpleFactorial(n -1)	
	}
}


func main() {
	result, _ := factorial(5)
	fmt.Println(result)

	_, err := factorial(-4)
	
	if err != nil {
		fmt.Println(err)
	}

	// Better way using uint
	result2 := simpleFactorial(5)
	fmt.Println(result2)

	var result1 uint
	// result1 = simpleFactorial(-3) //cannot do this
	fmt.Println(result1)

}
