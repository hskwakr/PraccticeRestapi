package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *echo.Echo) {
	app.GET("/books/", getBooks)
	app.GET("/book/:id", getBook)
	app.POST("/book/", newBook)
	app.DELETE("/book/:id", deleteBook)
}

func initDatabase() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func main() {
	app := echo.New()

	initDatabase()
	setupRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
