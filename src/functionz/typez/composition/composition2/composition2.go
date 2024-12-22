package main

import "fmt"

type sportsCar interface {
	turboON()
}


type luxury interface {
	autoPark()
}

type luxurySport interface {
	sportsCar
	luxury

	//can add new methods also... but then it will have to implement a function with receiver in the object
	// to implicit implement this interface
}

type bmw7 struct {}


func (b bmw7) turboON() {
	fmt.Println("Turbo ON")
}

func (b bmw7) autoPark() {
	fmt.Println("Automatic parking")
}

func main() {
	var b luxurySport = bmw7{}
	b.turboON()
	b.autoPark()
}