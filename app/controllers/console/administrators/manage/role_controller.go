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

	var dataCount int64
	if err := database.DB.Model(&models.Role{}).Count(&dataCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var data []models.Role
	if dataCount > 0 {
		if err := database.DB.Find(&data).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var dataResponse []RolesResponse = []RolesResponse{}
	for _, role := range data {
		dataResponse = append(dataResponse, RolesResponse{
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
			"list":  dataResponse,
			"total": dataCount,
		},
	})
}