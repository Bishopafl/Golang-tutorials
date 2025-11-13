package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(f) -> this prints &{0xc000120788}

	io.Copy(os.Stdout, f) // This displays the output of the file

}
