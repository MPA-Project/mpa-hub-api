package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/joho/godotenv/autoload"
	"myponyasia.com/hub-api/pkg/configs"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/middleware"
	"myponyasia.com/hub-api/pkg/routes"
	"myponyasia.com/hub-api/pkg/utils"
)

func main() {
	// Remove temporary files after 2 hours
	go utils.RemoveExpiredFiles()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Connect to the Database
	database.ConnectDB()

	// Storage config
	configs.S3Config()

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Monitor
	app.Get("/dashboard", monitor.New())

	// Base route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "Hi, welcome in My Pony Asia Hub API",
			"data": fiber.Map{
				"app": fiber.Map{
					"version": "v1.0.0",
					"build":   "12022022",
					"sha":     "07bf804cff6d5a10d2fac6bd56485ee00d25f5b56fd731f4f5f33944a7379b86",
				},
			},
		})
	})

	// Routes.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServer(app)

}
