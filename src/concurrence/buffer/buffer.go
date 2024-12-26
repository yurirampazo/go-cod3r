package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	fmt.Println("Executed!")
	ch <- 1
	fmt.Println("Executed 2!")
	ch <- 2
	fmt.Println("Executed 3!")
	ch <- 3
	fmt.Println("Executed 4!")
	ch <- 4
	fmt.Println("Executed 5!")
	ch <- 5
	fmt.Println("Executed 6!")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)

	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
