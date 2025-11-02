package main

import (
	"testConnectSQL/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.InitBooks()
	app.Get("/books", api.GetBooks)
	app.Get("/books/:id", api.GetBook)
	app.Post("/books/create", api.CreateBook)
	app.Put("/books/update/:id", api.UpdateBook)
	app.Delete("/books/delete/:id", api.DeleteBook)
	app.Post("/upload/", api.UploadFlie)
	app.Listen(":8000")
}

