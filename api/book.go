package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
var books []Book

func InitBooks() {
	books = append(books, Book{ID: 1, Title: "1984", Author: "George"})
	books = append(books, Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"})
}
func GetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusBadRequest)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {

	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookID {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusBadRequest)
}
func DeleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookDelete := new(Book)
	if err := c.BodyParser(bookDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...) // ... is used to flatten the slice
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func UploadFlie(c *fiber.Ctx) error{
	file , err := c.FormFile("image")
	if err != nil {
		return  c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	c.SaveFile(file, "./upload/" + file.Filename)			

	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("file uploaded successfully")
}