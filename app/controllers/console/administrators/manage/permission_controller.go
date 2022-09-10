package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
)

type PermissionResponse struct {
	ID   uuid.UUID
	Name string

	UserCount int
}

func PermissionList(c *fiber.Ctx) error {

	var dataCount int64
	if err := database.DB.Model(&models.Permission{}).Count(&dataCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var data []models.Permission
	if dataCount > 0 {
		if err := database.DB.Find(&data).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var dataResponse []PermissionResponse = []PermissionResponse{}
	for _, perm := range data {
		dataResponse = append(dataResponse, PermissionResponse{
			ID:        perm.ID,
			Name:      perm.Name,
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
