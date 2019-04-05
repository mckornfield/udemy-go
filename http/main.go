package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	fmt.Println(resp)

	// Crummy way

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Better way
	// io.Copy(os.Stdout, resp.Body)

	// Best way

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
