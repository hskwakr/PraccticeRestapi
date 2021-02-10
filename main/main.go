package main

import (
	"hskwakr/book/book"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/books/", book.GetBooks)
	app.GET("/book/:id", book.GetBook)
	app.POST("/book/", book.NewBook)
	app.DELETE("/book/:id", book.DeleteBook)

	app.Logger.Fatal(app.Start(":8080"))
}
