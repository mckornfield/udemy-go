package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the name of the file")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
