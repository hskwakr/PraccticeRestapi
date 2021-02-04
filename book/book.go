package book

import (
	"github.com/gofiber/fiber/v2"
)

// GetBooks gets all book.
func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

// GetBook gets a book.
func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Book")
}

// NewBook adds a new book.
func NewBook(c *fiber.Ctx) error {
	return c.SendString("Adds a new Book")
}

// DeleteBook deletes a book.
func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Deletes a Book")
}
