package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Devon", "Conway"}
	devon := person{firstName: "Devon", lastName: "Conway"}
	devon.firstName = "Devy"
	fmt.Printf("%+v", devon)

	taylor := person{
		firstName: "Ross",
		lastName:  "Taylor",
		contact: contactInfo{
			email:   "gagandeep@gmail.com",
			zipCode: 999008,
		},
	}
	fmt.Printf("%+v", taylor)
}
