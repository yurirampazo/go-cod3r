package main

import "fmt"

//Run all files in this package to see init power.

// go run *.go
func init() {
	fmt.Println("Initializing...")
}

func main() {
	fmt.Println("Main...")
}