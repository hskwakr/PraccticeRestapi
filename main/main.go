package main

import (
	"hskwakr/book/book"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.GET("/api/v1/book", book.GetBooks)
	app.GET("/app/v1/book/:id", book.GetBook)
	app.POST("/app/v1/book/", book.NewBook)
	app.DELETE("/app/v1/book/:id", book.DeleteBook)

	app.Logger.Fatal(app.Start(":1323"))
}
