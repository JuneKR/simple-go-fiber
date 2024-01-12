package main

import (
	"fmt"
)

type Book struct {
	ID     int    `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

var books []Book

func main() {
	books = append(books, Book{ID: 1, Title: "New Journey", Author: "Mike"})
	books = append(books, Book{ID: 2, Title: "New Journey", Author: "Mike"})

	fmt.Println(books)
}
