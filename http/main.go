package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	// if the error is not empty
	if err != nil {
		fmt.Println("Error: ", err) // print the error
		os.Exit(1)                  // exit
	}

	lw := logWriter{}

	// lw implements the write interface
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
