package main

type capsicumToppings struct {
	pizza pizza
}

func (c capsicumToppings) getPrice() int {
	return c.pizza.getPrice() + 10
}
