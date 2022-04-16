package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
)

type RolesResponse struct {
	ID    uuid.UUID
	Name  string
	Level int

	UserCount int
}

func RoleList(c *fiber.Ctx) error {

	var rolesCount int64
	if err := database.DB.Model(&models.Role{}).Count(&rolesCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var roles []models.Role
	if rolesCount > 0 {
		if err := database.DB.Find(&roles).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var rolesResponse []RolesResponse
	for _, role := range roles {
		rolesResponse = append(rolesResponse, RolesResponse{
			ID:        role.ID,
			Name:      role.Name,
			Level:     role.Level,
			UserCount: 0,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"list":  rolesResponse,
			"total": rolesCount,
		},
	})
}
