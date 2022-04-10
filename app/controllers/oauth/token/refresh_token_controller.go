package token

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils/authorization"
)

func RegenerateAccessToken(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	uuid := claims["uuid"].(string)

	var user models.User
	if err := database.DB.First(&user, "id = ?", uuid).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	at_token, rt_token, err := authorization.GenerateNewAccessToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"access_token":          at_token,
			"acces_token_expired":   time.Now().Add(15 * time.Minute),
			"refresh_token":         rt_token,
			"refresh_token_expired": time.Now().Add(365 * (24 * time.Hour)),
		},
	})
}
