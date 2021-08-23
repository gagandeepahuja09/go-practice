package main

func main() {
	receiption := &receiption{}
	doctor := &doctor{}
	chemist := &chemist{}
	cashier := &cashier{}
	receiption.setNext(doctor)
	doctor.setNext(chemist)
	chemist.setNext(cashier)

	abcPatient := &patient{name: "abc"}
	receiption.execute(abcPatient)
	chemist.execute(abcPatient)
}
