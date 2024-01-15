package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

type User struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

var books []Book

var memberUser = User{
	Email:    "user@example.com",
	Password: "password123",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
	})
}

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf(
		"URL = %s, Method = %s, Time = %s\n",
		c.OriginalURL(), c.Method(), start,
	)

	return c.Next()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "New Journey", Author: "Mike"})
	books = append(books, Book{ID: 2, Title: "New Journey 2", Author: "Mike"})

	app.Post("/login", login)

	app.Use(checkMiddleware)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/config", getEnv)

	app.Listen(":8081")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")
}

func getEnv(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "defaultsecret"
	}

	return c.JSON(fiber.Map{
		"SECRET": secret,
	})

}
