package main

import (
	"hskwakr/book/book"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *echo.Echo) {
	app.GET("/books/", book.GetBooks)
	app.GET("/book/:id", book.GetBook)
	app.POST("/book/", book.NewBook)
	app.DELETE("/book/:id", book.DeleteBook)
}

func initDatabase() {
	db, err := gorm.Open(mysql.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&book.Book{})
}

func main() {
	app := echo.New()

	setupRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
