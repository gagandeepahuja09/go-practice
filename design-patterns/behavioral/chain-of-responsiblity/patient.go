package main

type patient struct {
	name              string
	registrationDone  bool
	healthCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func (p *patient) refresh() {
	p.paymentDone = false
	p.medicineDone = false
	p.healthCheckUpDone = false
	p.registrationDone = false
}
