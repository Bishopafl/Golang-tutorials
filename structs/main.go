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
	peter := person{"Peter", "Parker"}
	fmt.Println(peter)
}

// 2. create a new value of type person
