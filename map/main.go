package main

import "fmt"

func main() {
	// literal syntax
	// key values get separated with a comma - like structs!
	// Can be used with other data types as well.

	colors := map[string]string{ // declaring a map with keys of type string and values of type string
		"red":   "#FF0000",
		"green": "#008000",
		"white": "#fffff", // last property always need a comma!
	}

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	// loop through each map
	// declaring, assigning the key values on :22
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
