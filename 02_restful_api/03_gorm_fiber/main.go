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
	defer database.DBConn.Close()
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(3000)

}
