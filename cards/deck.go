package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

// save a string to a file on local machine
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// obtain a new deck from a file
func newDeckFromFile(filename string) deck {
	// err is value of type 'error' - if nothing went wrong, value = nil
	bs, err := os.ReadFile(filename)

	// check if err is nul
	if err != nil {
		// Option #2 - log error and entirely quit program
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// string(bs) // Ace of Spades,Two of Spades,Three of Spades,...

	s := strings.Split(string(bs), ",")

	return deck(s) // want to get a deck, passing in the strings

}

// take the deck passed and randomize it
func (d deck) shuffle() {
	// pass the time object as a type for the source
	// this ensures that the number generated for randomization is 'now'
	// and will since we are using UnixNano - it will return a type int64
	source := rand.NewSource(time.Now().UnixNano()) // use as the source object as the basis for the random number generator
	r := rand.New(source)                           // assign 'r' the random number that was generated based on the NewSource() from the rand library
	for i := range d {
		newPosition := r.Intn(len(d) - 1) // get the random number based on one less of the length of the slice

		// swapping the position of i and new position of the card
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
