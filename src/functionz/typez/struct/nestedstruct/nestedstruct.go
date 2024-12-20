package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userId int
	itens  []item
}

func (o order) getTotalPrice() float64 {
	var total float64
	for _, it := range o.itens {
		total += it.price * float64(it.quantity)
	}

	return total
}

func main() {
	order := order{
		userId: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			{
				productID: 3,
				quantity:  100,
				price:     3.13,
			}, //redundant declaration
		},
	}

	fmt.Printf("Final Price $: %.2f", order.getTotalPrice())

}
