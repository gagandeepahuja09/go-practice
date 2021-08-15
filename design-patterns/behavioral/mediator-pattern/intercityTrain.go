package main

import "fmt"

type intercityTrain struct {
	mediator mediator
}

func (i *intercityTrain) requestArrival() {
	if i.mediator.canLand(i) {
		fmt.Println("Intercity train landing")
	} else {
		fmt.Println("Intercity train waiting")
	}
}

func (i *intercityTrain) departure() {
	fmt.Println("Intercity train leaving")
	i.mediator.notifyFree()
}

func (i *intercityTrain) permitArrival() {
	fmt.Println("Intercity train arrive permitted. Landing")
}
