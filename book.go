package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

// Book represents a book
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks gets all book.
func getBooks(c echo.Context) error {
	return c.String(http.StatusOK, "all books")
}

// GetBook gets a book.
func getBook(c echo.Context) error {
	return c.String(http.StatusOK, "A Single Book")
}

// NewBook adds a new book.
func newBook(c echo.Context) error {
	return c.String(http.StatusOK, "Adds a new Book")
}

// DeleteBook deletes a book.
func deleteBook(c echo.Context) error {
	return c.String(http.StatusOK, "Deletes a Book")
}
