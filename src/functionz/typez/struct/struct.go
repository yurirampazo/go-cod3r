package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// Method: function with receiver
func (p product) priceWithDiscount() float64 {
 return p.price * (1 - p.discount)
}


func round(number float64) string {
	return fmt.Sprintf("%.2f", number)
}


func main() {
	var product1 product
	product1 = product{
		name: "Pen",
		price: 1.79,
		discount: 0.05, 
	}

	fmt.Println(product1, round(product1.priceWithDiscount()))

	product2 := product{"Notebook", 1999.99, 0.10}
	fmt.Println(product2,  round(product2.priceWithDiscount()))

}
