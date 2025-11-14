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

	// the arrow acts to receive the value from the channel
	// this routine is put to sleep
	// ** Receiving messages from the Routine is blocking and does not continue. See slide #30

	// This prints each link message but is messy and unclean... will resolve shortly.
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // this makes our program hang because there isn't a seventh link...
}

// takes a link and checks
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// means something is wrong with the link
	if err != nil {
		fmt.Println(link, "might be down!")
		// send the information up to the channel to communicate to our main routine
		c <- "Might be down... I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yeah, it's up"

}
