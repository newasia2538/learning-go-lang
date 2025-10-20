package cards

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func HelloCards() {
	fmt.Println("Hello Card package")
}

func NewDeck() deck {
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

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) SaveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		errorLog(fmt.Sprintf("ERROR : %s", err))
		os.Exit(1)
	}
	stringFromByteSlice := strings.Split(string(bs), ", ")
	return deck(stringFromByteSlice)
}

func writeLog(filePath string, logMsg string) {
	os.WriteFile(filePath, []byte(logMsg), 0666)
}

func errorLog(logMsg string) {
	writeLog("logs/error.log", logMsg)
}
