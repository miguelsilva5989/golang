package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add new book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete a book")
}
