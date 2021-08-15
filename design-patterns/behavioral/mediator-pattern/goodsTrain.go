package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("Goods train landing")
	} else {
		fmt.Println("Goods train waiting")
	}
}

func (g *goodsTrain) departure() {
	fmt.Println("Goods train leaving")
	g.mediator.notifyFree()
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("Goods train arrive permitted. Landing")
}
