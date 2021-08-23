package main

import "fmt"

type receiption struct {
	next department
}

func (r *receiption) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Patient has already done registration ", p.name)
		r.next.execute(p)
		return
	}
	fmt.Println("Doing registration for patient ", p.name)
	p.registrationDone = true
	fmt.Println("Registration done for patient ", p.name)
	r.next.execute(p)
}

func (r *receiption) setNext(next department) {
	r.next = next
}
