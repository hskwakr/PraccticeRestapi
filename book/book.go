package book

import (
	"github.com/gofiber/fiber"
)

// GetBooks gets all book.
func GetBooks(c *fiber.Ctx) {
	c.SendString("All Books")
}

// GetBook gets a book.
func GetBook(c *fiber.Ctx) {
	c.SendString("A Single Book")
}

// NewBook adds a new book.
func NewBook(c *fiber.Ctx) {
	c.SendString("Adds a new Book")
}

// DeleteBook deletes a book.
func DeleteBook(c *fiber.Ctx) {
	c.SendString("Deletes a Book")
}
