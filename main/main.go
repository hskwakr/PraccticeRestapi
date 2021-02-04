package main

import (
	"hskwakr.practice/book/book"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/app/v1/book/:id", book.GetBook)
	app.Post("/app/v1/book/", book.NewBook)
	app.Delete("/app/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(": 3000")
}
