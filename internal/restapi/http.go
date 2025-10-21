package restapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InitializeHTTP() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !!")
	})

	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("HTTP starting port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func InitializeHTTPWithFiber() {
	app := fiber.New()

	app.Get("/helloFiber", func(c *fiber.Ctx) error {
		return c.SendString("Hello World Fiber!!")
	})

	fmt.Printf("Fiber starting port :8080")
	app.Listen(":8080")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Error Not Found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "HTTP Method is not allowed.", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello world !!!")
}
