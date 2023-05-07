package configuration

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/exception"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		BodyLimit:    10 * 1024 * 1024,
		ErrorHandler: exception.ErrorHandler,
	}
}
