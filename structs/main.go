package main

import "fmt"

// 1. tell Go what structs we will use

// ContactInfo struct
type contactInfo struct {
	email   string
	zipCode int
}

// person struct
type person struct {
	firstName   string
	lastName    string
	contactInfo // can take a custom type and use it in another custom type! Don't need to specify the field name - just the struct name
}

func main() {
	adam := person{
		firstName: "Adam", // NOTE: the only time a comma is used!!
		lastName:  "Lopez",
		contactInfo: contactInfo{ // embedded structs for reuse.
			zipCode: 33579,
			email:   "adam@test.com", // NOTE: all lines - even the last require commas!!
		},
	}

	adam.updateName("Noah") // did not update
	adam.print()
}

// update the name of the person struct
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// print information about the person
// by taking a person as a receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
