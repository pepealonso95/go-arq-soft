package routes

import (
	"c7/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, controller *controllers.BookController) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/books", controller.GetBooks)      // get list of all books
	route.Get("/book/:title", controller.GetBook) // get one book by ID

	// Routes for POST method:
	route.Post("/books", controller.AddBook) // register a new user
}
