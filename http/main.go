package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	// if the error is not empty
	if err != nil {
		fmt.Println("Error: ", err) // print the error
		os.Exit(1)                  // exit
	}

	// print the line out from the request
	fmt.Println(resp)
}
