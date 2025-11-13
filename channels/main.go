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

	// when ignoring the 'key' and only wanting the value,
	// put an underscore where the key should be
	for _, link := range links {
		checkLink(link)
	}

}

// takes a link and checks
func checkLink(link string) {
	_, err := http.Get(link)
	// means something is wrong with the link
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")

}
