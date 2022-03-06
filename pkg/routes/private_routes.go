package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth/token"
	"myponyasia.com/hub-api/app/controllers/v1/user"
	"myponyasia.com/hub-api/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(app *fiber.App) {
	// Create user routes group.
	routeUser := app.Group("/v1/user/me", middleware.JWTSessionProtected())
	routeUser.Get("/", user.Me)

	// Group oauth routes
	routeOauth := app.Group("/oauth")
	routeOauth.Post("refresh-token", middleware.JWTRefreshProtected(), token.RegenerateAccessToken)
}
