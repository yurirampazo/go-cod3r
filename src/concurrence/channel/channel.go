package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // Sending data to a channel (Write)
	<-ch  // Getting data from a channel (Read)

	ch <- 2
	fmt.Println(<-ch)





}