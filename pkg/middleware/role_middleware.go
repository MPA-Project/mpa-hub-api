package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RoleAdmin(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	data := claims["data"].(map[string]interface{})
	roles := data["roles"].([]interface{})

	isFound := false
	for _, value := range roles {
		if value == "ADMIN" {
			isFound = true
		}
	}

	if !isFound {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "You are not authorized to perform this action",
		})
	}

	return c.Next()
}

func RoleModerator(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	data := claims["data"].(map[string]interface{})
	roles := data["roles"].([]interface{})

	isFound := false
	for _, value := range roles {
		if value == "MODERATOR" {
			isFound = true
		}
	}

	if !isFound {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "You are not authorized to perform this action",
		})
	}

	return c.Next()
}
