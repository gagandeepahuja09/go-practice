package main

import "fmt"

func main() {
	cards := deck{newCard(), "Five of Diamonds"}

	cards.print()
	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
