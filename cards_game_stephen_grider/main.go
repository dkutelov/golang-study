package main

import "fmt"

func main() {
	cards := []string {"Ace of Spades", newCard()}
	card := newCard()
	// append does not modifies the existing slice but returns a new one
	cards = append(cards, card)
	fmt.Println(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}