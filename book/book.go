package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetBooks gets all book.
func GetBooks(c echo.Context) error {
	return c.String(http.StatusOK, "all books")
}

// GetBook gets a book.
func GetBook(c echo.Context) error {
	return c.String(http.StatusOK, "A Single Book")
}

// NewBook adds a new book.
func NewBook(c echo.Context) error {
	return c.String(http.StatusOK, "Adds a new Book")
}

// DeleteBook deletes a book.
func DeleteBook(c echo.Context) error {
	return c.String(http.StatusOK, "Deletes a Book")
}
