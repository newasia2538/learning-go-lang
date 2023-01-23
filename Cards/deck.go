package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	deck := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+" of "+suit)
		}
	}

	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		errorLog(fmt.Sprintf("ERROR : %s", err))
		os.Exit(1)
	}
	stringFromByteSlice := strings.Split(string(bs), ", ")
	return deck(stringFromByteSlice)
}

func writeLog(filePath string, logMsg string) {
	ioutil.WriteFile(filePath, []byte(logMsg), 0666)
}

func errorLog(logMsg string) {
	writeLog("logs/error.log", logMsg)
}
