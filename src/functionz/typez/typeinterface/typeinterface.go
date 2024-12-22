package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)


	thing = 3
	fmt.Println(thing)

	thing = "hey"
	fmt.Println(thing)


	type dynamic interface{}

	var thing2 dynamic = "Hi"
  fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = course{"Golang: Exploring this amazing lenguage"}
	fmt.Println(thing2)
}
