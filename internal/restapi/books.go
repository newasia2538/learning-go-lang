package restapi

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/newasia2538/learning-go-lang/internal/middleware"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	_ "github.com/newasia2538/learning-go-lang/docs"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func InitializeBooksAPI() {
	fmt.Println(os.Getwd())
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env file failed")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, HEAD",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	books = append(books, Book{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"})
	books = append(books, Book{ID: 2, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin"})

	fmt.Println(books)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Post("/login", login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
	}))

	app.Use(middleware.CheckMiddleware)

	app.Get("/", getHTML)
	app.Get("/api/config", getAPIConfig)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/book", addBook)
	app.Put("/book/:id", updateBook)
	app.Delete("/book/:id", deleteBook)
	app.Post("/upload", uploadFile)

	app.Listen(":8181")
}

// GetBook godoc
// @Summary      get list of all books
// @Description  get list of books
// @Tags         books
// @Accept       json
// @Produce      json
// @security	 ApiKeyAuth
// @Success      200  {array}  Book
// @Router       /books [get]
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func addBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.SendStatus(fiber.StatusCreated)
}

func updateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == id {
			books[index].Title = bookUpdate.Title
			books[index].Author = bookUpdate.Author
			return c.Status(fiber.StatusOK).JSON(books[index])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			return c.Status(fiber.StatusNoContent).SendString(fmt.Sprintf("book ID '%d' was deleted", id))
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./resources/images/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload success !")
}

func getHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello World !",
	})
}

func getAPIConfig(c *fiber.Ctx) error {
	// if value, exists := os.LookupEnv("SECRET"); exists {
	// 	return c.JSON(fiber.Map{
	// 		"SECRET": value,
	// 	})
	// }

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
