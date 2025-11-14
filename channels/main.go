package main

import (
	"fmt"
	"net/http"
	"time"
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

	// lets clean it up
	// when we receive a message through channel c
	// the new value is assigned to 'l'
	//
	// we pass 'l' to the function literal
	// that string is copied in memory
	// then the Go Routine has access to the copy
	// as opposed to the original value of 'l'
	//
	// 'l' can change as much as it pleases, we don't have to worry about Go routine accessing that same address in memory.
	for l := range c {
		// we are going to put a function literal here
		// to make sure that we are only pausing the function after the routines return a value
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // NEED TO HAVE AN EXTRA SET OF PARATHESIS TO ENVOKE THE FUNCTION LITERAL
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
