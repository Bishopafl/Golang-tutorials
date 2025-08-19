package main

import "fmt"

func main() {
	// literal syntax
	// key values get separated with a comma - like structs!
	// Can be used with other data types as well.

	// colors := map[string]string{ // declaring a map with keys of type string and values of type string
	// 	"red":   "#FF0000",
	// 	"green": "#008000",
	// }

	// var colors map[string]string // usually take this approach to see what we want to put into a map

	colors := make(map[string]string) // makes a new map of key/values of strings ** we can add values to these if needed later

	colors["white"] = "#fffff" // in maps the keys are typed, and need to be of the appropriate type

	fmt.Println(colors)
}
