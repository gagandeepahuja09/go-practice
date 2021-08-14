package main

import "fmt"

func main() {
	peppyPaneerWithOnionToppings := onionToppings{
		pizza: peppyPaneer{},
	}
	fmt.Println(peppyPaneerWithOnionToppings.getPrice())

	spicyChickerWithCapAndOnToppings := capsicumToppings{
		pizza: onionToppings{
			pizza: spicyChicken{},
		},
	}
	fmt.Println(spicyChickerWithCapAndOnToppings.getPrice())
}
