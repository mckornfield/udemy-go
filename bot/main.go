package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (b englishBot) getGreeting() string {
	// Custom logic for english
	return "Hello There"
}

func (b spanishBot) getGreeting() string {
	// Custom logic for spanish
	return "Hola"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}
	printGreeting(englishBot)
	printGreeting(spanishBot)
}
