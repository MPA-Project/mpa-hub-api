package routes

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/controllers/oauth/account"
	// v1 "myponyasia.com/hub-api/app/controllers/v1"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {

	// Group oauth routes
	routeOauth := app.Group("/oauth")
	routeOauth.Post("signin", account.Login)
	routeOauth.Post("signup", account.Register)
	routeOauth.Post("token-verify", account.RequestTokenVerify)
	routeOauth.Post("forgot-password", account.ForgotPassword)
	routeOauth.Post("forgot-password-confirm", account.ForgotPasswordConfirm)
	routeOauth.Post("email-verification", account.EmailVerification)
}
