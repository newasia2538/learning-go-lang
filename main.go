package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/newasia2538/learning-go-lang/cards"
)

func main() {
	fmt.Println("Yoooo !!")
	uuid := uuid.New()
	fmt.Println("UUID : ", uuid)
	cards.HelloCards()
}
