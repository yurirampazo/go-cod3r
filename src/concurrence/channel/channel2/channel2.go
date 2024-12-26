package main

import (
	"fmt"
	"time"
)

// Channel (canal) -> Comunication between goroutines
// channel is a type


func twoThreeFourTimes(base int, c chan int) {
  time.Sleep(time.Second)
  c <- 2 * base //sending data to channel

  time.Sleep(time.Second)
  c <- 3 * base

  time.Sleep(3 * time.Second)
  c <- 4 * base
}

func main() {
  c := make(chan int)

  go twoThreeFourTimes(2, c)
  fmt.Println("A")

  a, b := <-c, <-c //getting data from channel, first 2 * base then 3 * base
  fmt.Println("B")
  fmt.Println(a, b)
  fmt.Println(<-c)


}