package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
)

func Me(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"id":    user.ID.String(),
			"name":  user.Username,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
