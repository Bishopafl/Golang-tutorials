package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	// both are of type deck - any value with type of deck can use this function
	hand.print()
	remainingCards.print()

}
