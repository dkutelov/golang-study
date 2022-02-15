package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toSting() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toSting()), 0666) //anyone can read and write
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		
		os.Exit(1) //Abort program
	}

	return strings.Split(string(bs), ",")
}