package main

import (
	"fmt"
	"io"
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

	io.Copy(os.Stdout, resp.Body)
}
