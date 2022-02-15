package main

func main() {
	// append does not modifies the existing slice but returns a new one
	// cards = append(cards, card)
	cards := newDeck()
	cards.saveToFile("mycards")
	
	//hand, ramainingCards := deal(cards, 5)
	//cards.print()

	// hand.print()
	// fmt.Println("---")
	// ramainingCards.print()

	//greeting := "Hi there!"
	// []byte(greeting) //convert string to byte slice
}

// func newCard() string {
// 	return "Five of Diamonds"
// }