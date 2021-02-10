package main

import (
	"hskwakr/book/book"

	"github.com/labstack/echo/v4"
)

func setupRoutes(app *echo.Echo) {
	app.GET("/books/", book.GetBooks)
	app.GET("/book/:id", book.GetBook)
	app.POST("/book/", book.NewBook)
	app.DELETE("/book/:id", book.DeleteBook)
}

func main() {
	app := echo.New()

	setupRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
