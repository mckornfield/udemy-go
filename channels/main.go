package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {

	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down!", error)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}