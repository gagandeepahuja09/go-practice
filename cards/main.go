package main

import "fmt"

// we can't do shorthand variable initialization outside main
// because outside main variable can only be declared & not initialized

func main() {
	// var card string = "Ace of Spades"
	// this is shorthand syntax + we can let go infer the type on its own.
	// := is shorthand syntax for initialization.
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)
}
