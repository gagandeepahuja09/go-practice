package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) requestArrival() {
	if p.mediator.canLand(p) {
		fmt.Println("Passenger train landing")
	} else {
		fmt.Println("Passenger train waiting")
	}
}

func (p *passengerTrain) departure() {
	fmt.Println("Passenger train leaving")
	p.mediator.notifyFree()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("Passenger train arrive permitted. Landing")
}
