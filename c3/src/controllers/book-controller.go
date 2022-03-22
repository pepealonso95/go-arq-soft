package controllers

import (
	"c3/src/models"
	"c3/src/repositories"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	repo *repositories.BookRepo
}

// Esto es un Factory Method porque no hay constructores
// En este caso lo usamos para iniciarlizar la variable privada
func NewBookController(repo *repositories.BookRepo) *BookController {
	return &BookController{repo: repo}
}

// Endpoint para traer los libros
func (controller *BookController) GetBooks(c *fiber.Ctx) error {

	// Get all books.
	books, err := controller.repo.GetBooks()
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})
}

func (controller *BookController) GetBook(c *fiber.Ctx) error {

	// Get book by ID.
	book, err := controller.repo.GetBook(c.Params("title"))
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"book":  nil,
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}

func (controller *BookController) AddBook(c *fiber.Ctx) error {
	var book models.Book
	// Tengo que pasar la direccion de memoria asi va a pone las variables
	// Que vienen en el contexto declaradas segun el json del modelo
	fmt.Println(string(c.Body()))

	book.CreatedAt = time.Now()
	err := json.Unmarshal(c.Body(), &book)
	if err != nil {
		// Return, if book has invalid fields
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"book":  nil,
		})
	}

	// Add new book
	err = controller.repo.AddBook(book)
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"book":  nil,
		})
	}

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Book created succesfully!",
		"book":  book,
	})
}
