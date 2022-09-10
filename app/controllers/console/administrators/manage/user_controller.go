package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
)

type UserResponse struct {
	ID       uuid.UUID
	Username string
	Email    string
}

func UserList(c *fiber.Ctx) error {

	var dataCount int64
	if err := database.DB.Model(&models.User{}).Count(&dataCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var data []models.User
	if dataCount > 0 {
		if err := database.DB.Find(&data).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var dataResponse []UserResponse = []UserResponse{}
	for _, user := range data {
		dataResponse = append(dataResponse, UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
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
