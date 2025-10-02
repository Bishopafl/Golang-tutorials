package main

import "fmt"

// adding the interface type
type bot interface {
	getGreeting() string // takes in the getGreeting function and identifies it returns a string
}

// creating empty structs that I can set receivers up for
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// custom spanish greeting
	return "Hola!"
}
