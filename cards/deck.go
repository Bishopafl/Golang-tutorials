package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// Create an empty deck
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// underscores are used in go to tell it that there is a variable there but we don't care about it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d is a receiver on a function and is similar to 'this' or 'self'
// Basically - any variable of type "deck"
//		now gets access to the "print" method
//
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
