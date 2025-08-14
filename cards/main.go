package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	// Important Note:
	// takes the cards slice, add a new string to end - append() does not modify existing slice,
	// instead it returns a new slice assigned to the variable cards
	cards = append(cards, "Six of Hearts")

	// How to iterate over cards:
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Ace of Spades"
}
