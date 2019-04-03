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

	fmt.Println(hand.toString())
	hand.saveToFile("my_cards")
	hand2 := newDeckFromFile("my_cards")
	hand2.print()
}
