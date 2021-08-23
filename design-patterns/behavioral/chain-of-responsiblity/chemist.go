package main

import "fmt"

type chemist struct {
	next department
}

func (c *chemist) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicenes already bought for patient ", p.name)
		c.next.execute(p)
		return
	}
	fmt.Println("Buying medicenes for patient ", p.name)
	p.medicineDone = true
	fmt.Println("Medicenes bought for patient ", p.name)
	c.next.execute(p)
}

func (c *chemist) setNext(next department) {
	c.next = next
}
