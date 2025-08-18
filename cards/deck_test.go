package main

import "testing"

const DeckLength = 16
const FirstCardValue = "Ace of Spades"
const LastCardValue = "Four of Clubs"

// Notice the capital for the func name here...
// t stands for 'test handler'
/*
Test Notes:
	Go doesn't know we have three different tests running in this file.
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != DeckLength {
		t.Errorf("Excected deck length of %v but got %v", DeckLength, len(d))
	}

	if d[0] != FirstCardValue {
		t.Errorf("Expected first card of %v but got %v", FirstCardValue, d[0])
	}

	if d[len(d)-1] != LastCardValue {
		t.Errorf("Expected last card of %v but got %v", LastCardValue, d[len(d)-1])
	}

}
