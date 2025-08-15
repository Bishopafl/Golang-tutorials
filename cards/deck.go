package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// Create an empty deck
	cards := deck{}

	// Define and initilize card variables
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// underscores are used in go to tell it that there is a variable there but we don't care about it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself.
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d is a receiver on a function and is similar to 'this' or 'self'
// Basically - any variable of type "deck"
//
//	now gets access to the "print" method
//
// this prints to the line the deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal functionality
// two separate return values in return.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Turn deck into a string
// Need another library to perform this
func (d deck) toString() string {
	// takes a slice of string, joins it and removes the comma in the slice
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
