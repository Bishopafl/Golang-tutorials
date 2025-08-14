package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// String example:
func newCard() string {
	return "Ace of Spades"
}

// Float example:
// func newCard() float64 {
// 	return 3.14
// }

// Integer example:
// func newCard() int {
// 	return 314
// }
