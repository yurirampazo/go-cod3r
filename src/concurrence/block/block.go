package main

import (
	"fmt"
	"time"
)


func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // blocking action
	fmt.Println("Only after read")

	// c <- 2 //test -> uncoment for threat deadlock
	// fmt.Println("ReallyEndend")
}


func main() {
	c := make(chan int) // channel without buffer

	go routine(c)

	fmt.Println(<-c) // blocking operation
	fmt.Println("Read")
	fmt.Println(<-c) //deadlock
	fmt.Println("End")


}