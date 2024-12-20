package main

import "fmt"

type car struct {
	name string
	actualSpeed int
}

type ferrari struct {
	car // anonymous field
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.actualSpeed = 0
	f.turboOn = true


	fmt.Printf("La Ferrari %s is with turbo on? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}