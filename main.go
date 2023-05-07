package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	_ "github.com/joho/godotenv/autoload"
	"myponyasia.com/hub-api/common"
	"myponyasia.com/hub-api/configuration"
	_ "myponyasia.com/hub-api/docs"
	"myponyasia.com/hub-api/exception"
)

// @title MPA HUB API
// @version 1.0.0
// @description MPA HUB API Documentation
// @termsOfService http://myponyasia.com/terms/
// @contact.name API Support
// @contact.email support@myponyasia.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {

	// Remove temporary files after 2 hours
	go common.RemoveExpiredFiles()

	//setup fiber
	app := fiber.New(configuration.FiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())

	// // Monitor
	// app.Get("/dashboard", monitor.New())

	// // Base route
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"status":  true,
	// 		"message": "Hi, welcome in My Pony Asia Hub API",
	// 		"data": fiber.Map{
	// 			"app": fiber.Map{
	// 				"version": "v1.0.0",
	// 				"build":   "12022022",
	// 				"sha":     "07bf804cff6d5a10d2fac6bd56485ee00d25f5b56fd731f4f5f33944a7379b86",
	// 			},
	// 		},
	// 	})
	// })

	// // Routes.
	// routes.PublicRoutes(app)  // Register a public routes for app.
	// routes.PrivateRoutes(app) // Register a private routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//start app
	err := app.Listen(os.Getenv("SERVER.PORT"))
	exception.PanicLogging(err)

}
