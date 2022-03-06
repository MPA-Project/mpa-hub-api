package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins: "http://localhost:4200, https://oauth.myponyasia.com, https://myponyasia.com",
		}),

		// Add simple logger.
		logger.New(),

		// Rate limiter
		limiter.New(limiter.Config{
			// Next: func(c *fiber.Ctx) bool {
			// 	return c.IP() == "127.0.0.1"
			// },
			Max:        60,
			Expiration: 60 * time.Second,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"error": true,
					"msg":   "Too Many Requests",
				})
			},
		}),

		// Etag
		etag.New(etag.Config{
			Weak: true,
		}),

		// Favicon
		favicon.New(favicon.Config{
			File: "./public/favicon.ico",
		}),

		// Request-id
		requestid.New(requestid.Config{
			Header: "X-Request-Id",
		}),
	)
}
