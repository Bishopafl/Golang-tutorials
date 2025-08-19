package main

import "fmt"

func main() {
	// literal syntax
	// key values get separated with a comma - like structs!
	colors := map[string]string{ // declaring a map with keys of type string and values of type string
		"red":   "#FF0000",
		"green": "#008000",
	}

	fmt.Println(colors)
}
