package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth/account"
	v1 "myponyasia.com/hub-api/app/controllers/v1"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {

	// Group V1 routes
	routeV1 := app.Group("/v1")
	routeV1.Get("/users", v1.GetUsers) // Get all users
	// routeV1.Get("/user/:userId", v1.GetUser) // get one user by ID
	// routeV1.Post("/user", v1.CreateUser)     // create a new user

	// Group oauth routes
	routeOauth := app.Group("/oauth")
	routeOauth.Post("signin", account.Login)
	routeOauth.Post("signup", account.Register)

	// route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
