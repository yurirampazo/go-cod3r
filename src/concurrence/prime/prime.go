package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int, c chan int) {
	begin := 2
	for i := 0; i < n; i++ {
		for prime := begin; ; prime++ {
			if isPrime(prime) {
				c <- prime
				begin = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primes(cap(c), c)

	for prime := range c {
		fmt.Printf("%d ", prime)
	}
	fmt.Println("END")
}
