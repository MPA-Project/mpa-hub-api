package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth/account"
	"myponyasia.com/hub-api/app/controllers/oauth/token"
	"myponyasia.com/hub-api/app/controllers/utils/uploads"
	"myponyasia.com/hub-api/app/controllers/v1/user"
	"myponyasia.com/hub-api/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(app *fiber.App) {

	// Group oauth routes
	routeOauth := app.Group("/oauth")
	routeOauth.Post("refresh-token", middleware.JWTRefreshProtected(), token.RegenerateAccessToken)
	routeOauth.Post("roles", middleware.JWTSessionProtected(), account.Roles)
	routeOauth.Post("permissions", middleware.JWTSessionProtected(), account.Permissions)

	// Group user routes
	routeUser := app.Group("/v1/user/me", middleware.JWTSessionProtected())
	routeUser.Get("/", user.Me)

	// Group utilities routes
	routeUtils := app.Group("/utils")
	// routeUtils.Post("/upload-temporary", middleware.JWTSessionProtected(), uploads.UploadTemporary)
	routeUtils.Post("/upload-temporary", uploads.UploadTemporary)
}
