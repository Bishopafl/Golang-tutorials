package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		// need to ensure the protocol is there within the links or it will not be read correctly
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://adamlopez.com",
	}

	// Channel setup!
	// 	Treat this value like any other value in the application
	c := make(chan string)

	// when ignoring the 'key' and only wanting the value,
	// put an underscore where the key should be
	for _, link := range links {
		// the `go` keyword creates a 'Go Routine' inside of our program - think of it that runs code line by line
		// **ALWAYS DO THIS BEFORE A FUNCTION**
		go checkLink(link, c)

	}

	// C inspired for loop
	// still waiting for a message to come through the channel until the next iteration of the loop occurs

	// Might look a little hackish - but many methods in Go take this approach
	for {
		// INFINITE LOOP!
		// will resolve when there are no more values to receive through the channel
		go checkLink(<-c, c) // receieve of the value through the channel
	}

}

// takes a link and checks
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// means something is wrong with the link
	if err != nil {
		fmt.Println(link, "might be down!")
		// send the information up to the channel to communicate to our main routine
		c <- link // rather than pushing a fixed string, we will pass the link or URL
		return
	}
	fmt.Println(link, "is up!")
	c <- link // rather than pushing a fixed string, we will pass the link or URL

}
