package main

// declare variable out of function
var c string

func main() {
	// declare variable and initialize value
	deck := newDeck()
	deck.saveToFile("my-cards")
}
