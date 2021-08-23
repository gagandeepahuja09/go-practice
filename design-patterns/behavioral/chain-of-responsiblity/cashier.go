package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment already done for patient ", p.name)
		return
	}
	fmt.Println("Completing payment for patient ", p.name)
	p.paymentDone = true
	fmt.Println("Payment completed for patient ", p.name)
}

func (c *cashier) setNext(next department) {
	c.next = next
}
