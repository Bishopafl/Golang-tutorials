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

	// declare byte slice
	// bs := []byte{}

	// built in func that taks a byte and number of empty slices the slice can be made with
	bs := make([]byte, 99999) // give me a byte slice that has n number of slices in it

	// Read function is going to read data into the slice until it's full.
	// Read function will not be able to read data into
	// 99999 is an assumption of the amount of slices injected into reader interface when http.get is called.
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
