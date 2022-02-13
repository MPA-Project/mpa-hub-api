package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {

	// Group Oauth routes
	routeOauth := app.Group("/oauth/")
	routeOauth.Get("/users", oauth.GetUsers)       // Get all users
	routeOauth.Get("/user/:userId", oauth.GetUser) // get one user by ID
	routeOauth.Post("/user", oauth.CreateUser)     // create a new user

	// route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
