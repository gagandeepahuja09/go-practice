package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Devon", "Conway"}
	devon := person{firstName: "Devon", lastName: "Conway"}
	pointerToDevon := &devon
	pointerToDevon.updateName("Devy")
	fmt.Printf("%v\n", &pointerToDevon)

	taylor := person{
		firstName: "Ross",
		lastName:  "Taylor",
		contactInfo: contactInfo{
			email:   "gagandeep@gmail.com",
			zipCode: 999008,
		},
	}

	devon.print()
	taylor.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
