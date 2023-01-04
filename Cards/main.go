package main

import (
	"fmt"
)

var c string

func main() {
	fmt.Println("CARDS !!!!")
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
