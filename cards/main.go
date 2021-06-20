package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
