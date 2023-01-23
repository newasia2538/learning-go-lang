package main

// declare variable out of function
var c string

func main() {
	// declare variable and initialize value
	deck := newDeckFromFile("my-cards")
	deck.print()
}
