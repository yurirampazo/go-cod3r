package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}


/**
	Interfaces are implemented in implicit way
*/
func (p person) toString() string {
	return p.name + " " + p.lastname
}

func(p product) toString() string {
	return fmt.Sprintf("%s - $%.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}


func main() {
	var thing printable = person{"yuri", "legal"}
	fmt.Println(thing.toString())
	print(thing)
	fmt.Println(thing)
	
	thing = product{"notebook", 6000.99}
	fmt.Println(thing.toString())
	print(thing)
	fmt.Println(thing)
	

	p2 := product{"T-Shirt", 30}
	print(p2)
	

}
