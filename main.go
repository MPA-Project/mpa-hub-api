package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "Hi, welcome in My Pony Asia API",
			"data": fiber.Map{
				"app": fiber.Map{
					"version": "v1.0.0",
					"build":   "11022022",
					"sha":     "07bf804cff6d5a10d2fac6bd56485ee00d25f5b56fd731f4f5f33944a7379b86",
				},
			},
		})
	})

	app.Listen(":3000")
}
