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

	// os.Stdout implements the Writer interface
	// second param resp.Body implements the Reader interface
	io.Copy(os.Stdout, resp.Body)
}
