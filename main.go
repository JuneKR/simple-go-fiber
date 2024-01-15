package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "New Journey", Author: "Mike"})
	books = append(books, Book{ID: 2, Title: "New Journey 2", Author: "Mike"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("books/:id", deleteBook)

	app.Listen(":8081")
}
