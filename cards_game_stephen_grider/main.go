package main

import "fmt"

func main() {
	// cards := deck {"Ace of Spades", newCard()}
	// card := newCard()
	// append does not modifies the existing slice but returns a new one
	// cards = append(cards, card)
	cards := newDeck()
	hand, ramainingCards := deal(cards, 5)
	//cards.print()

	hand.print()
	ramainingCards.print()
	fmt.Println("---")
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }