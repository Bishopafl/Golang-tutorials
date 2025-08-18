package main

import "fmt"

// 1. tell go what fields the person struct has
// 		- custom type for our project

type person struct {
	firstName string
	lastName  string
}

func main() {
	// every person has a first name and last name
	// first and last values are assumed in Go
	// that they follow the struct built on the
	// order of the fields.

	/*
		Two ways to declare a new struct in Go
	*/
	// first way:
	// peter := person{"Peter", "Parker"}

	// second way:
	var peter person // one way to declare a new struct in go
	peter.firstName = "Peter"
	peter.lastName = "Parker"

	fmt.Println(peter)
	fmt.Printf("%+v", peter)
}

// 2. create a new value of type person
