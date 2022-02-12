package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/user", oauth.CreateUser) // create a new user

	// Routes for GET method:
	// route.Get("/books", controllers.GetBooks)              // get list of all books
	// route.Get("/book/:id", controllers.GetBook)            // get one book by ID
	// route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
