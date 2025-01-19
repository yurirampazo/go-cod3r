package main

import (
	"fmt"
	"reflect"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := person{name:"Peter", lastname: "Parker"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Tony Stark")
	fmt.Println(p1.getFullName())
	fmt.Println()


	p2 := &p1
	p3 := *&p1
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println()
	fmt.Println(p2.getFullName())
	fmt.Println(p3.getFullName())
	fmt.Println()
	
	p1.setFullName("Max Verstappen")
	fmt.Println(p1.getFullName())
	fmt.Println(p2.getFullName())
	fmt.Println(p3.getFullName())
	fmt.Println()

	if p1 == p1 {
		fmt.Println("equals")
		fmt.Println(reflect.TypeOf(p1))
		fmt.Println(reflect.TypeOf(p2))
	} else {
		fmt.Println("not equals")
		fmt.Println(reflect.TypeOf(p1))
		fmt.Println(reflect.TypeOf(p2))
	}

	if p1 == p3 {
		fmt.Println("equals")
		fmt.Println(reflect.TypeOf(p1))
		fmt.Println(reflect.TypeOf(p3))
	} else {
		fmt.Println("not equals")
		fmt.Println(reflect.TypeOf(p1))
		fmt.Println(reflect.TypeOf(p3))
	}

}
