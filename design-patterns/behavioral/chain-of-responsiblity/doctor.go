package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Patient has already done health checkup ", p.name)
		d.next.execute(p)
		return
	}
	fmt.Println("Doing health checkup for patient ", p.name)
	p.healthCheckUpDone = true
	fmt.Println("Health checkup done for patient ", p.name)
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}
