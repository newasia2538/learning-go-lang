package rest_api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func InitializeBooksAPI() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"})
	books = append(books, Book{ID: 2, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin"})

	fmt.Println(books)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/book", addBook)
	app.Put("/book/:id", updateBook)
	app.Delete("/book/:id", deleteBook)

	app.Listen(":8181")
}

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
