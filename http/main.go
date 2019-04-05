package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
	io.Copy(os.Stdout, resp.Body)
}
