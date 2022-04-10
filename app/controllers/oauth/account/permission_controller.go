package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Permissions(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	data := claims["data"].(map[string]interface{})
	permissions := data["permissions"].([]interface{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data":    permissions,
	})
}
