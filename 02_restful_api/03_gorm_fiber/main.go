// https://www.youtube.com/watch?v=Iq2qT0fRhAA

package main

import (
	"fmt"

	"./book"
	"./database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)

}
