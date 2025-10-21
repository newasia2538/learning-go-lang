package main

import (
	"fmt"

	"github.com/newasia2538/learning-go-lang/internal/restapi"
)

// @title Books Example API
// @version 1.0
// @description This is a sample swagger for Books API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8181
// @BasePath /
func main() {
	// fmt.Println("Yoooo !!")
	// uuid := uuid.New()
	// fmt.Println("UUID : ", uuid)
	// cards.HelloCards()
	fmt.Println("========================================")
	restapi.InitializeBooksAPI()
}
