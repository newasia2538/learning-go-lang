package main

import (
	"fmt"
)

// declare variable out of function
var c string

func main() {
	// declare variable and initialize value
	card := newCard()
	// declare slice of string type and initailize value
	cards := []string{"Ace of Diamonds", newCard()}
	deck := deck{"Ace of Diamonds", newCard()}
	// add new record into slice
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println("--------------------------------")
	newDeck := newDeck()
	newDeck.print()
	hand, remainingDeck := deal(newDeck, 5)
	fmt.Println("--------------------------------")
	deck.print()
	fmt.Println("--------------------------------")
	fmt.Println(card)
	fmt.Println("--------------------------------")
	hand.print()
	fmt.Println("--------------------------------")
	remainingDeck.print()
	fmt.Println("--------------------------------")
	fmt.Println(newDeck.toString())

}

// declare function return string type value
func newCard() string {
	return "Five of Diamonds"
}
