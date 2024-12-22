package main

import "fmt"

type sportsCar interface {
	turnTurboOn()
}

type ferrari struct {
	model       string
	turboOn     bool
	actualSpeed int
}


func (f *ferrari) turnTurboOn() {
	f.turboOn = true
}

func main() {

	car1 := ferrari{"f40", false, 0}
	car1.turnTurboOn()

	var car2 sportsCar = &ferrari{"f40", false, 0}
	car2.turnTurboOn()

	fmt.Println(car1, car2)


}
