package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Devon", "Conway"}
	alex := person{firstName: "Devon", lastName: "Conway"}
	alex.firstName = "Devy"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
