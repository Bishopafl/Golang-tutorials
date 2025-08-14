package main

func main() {

	// Because cards is of type deck - we can call the print() method on it.
	cards := deck{"Ace of Diamonds", newCard()}

	// Important Note:
	// takes the cards slice, add a new string to end - append() does not modify existing slice,
	// instead it returns a new slice assigned to the variable cards
	cards = append(cards, "Six of Hearts")

	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}
