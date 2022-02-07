package main

import "fmt"

// Create new type of 'deck'

type deck []string //extends slice of string

// new receiver func of deck type
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}