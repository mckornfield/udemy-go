package main

import "fmt"

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	cards.print()
	hand, remainingCards := deal(cards, 7)
	fmt.Println("Remaining cards are")
	remainingCards.print()
	fmt.Println("Hand is")
	hand.print()
}
