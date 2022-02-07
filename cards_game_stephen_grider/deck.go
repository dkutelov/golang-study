package main

import "fmt"

// Create new type of 'deck'

type deck []string //extends slice of string

// new receiver func of deck type - d is like self
// By convention we refer to receiver with 1 or 2 letters from the type
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck {}
	cardSuites := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string {"Ace", "Two", "Three", "Four"}
	
	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardsValues {
			newCard := cardValue + " of " + cardSuite
			cards = append(cards, newCard)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck){
	return d[:handsize], d[handsize:]
}