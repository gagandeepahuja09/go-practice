package main

type onionToppings struct {
	pizza pizza
}

func (o onionToppings) getPrice() int {
	return o.pizza.getPrice() + 7
}
