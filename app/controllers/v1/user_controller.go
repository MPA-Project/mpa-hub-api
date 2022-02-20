package v1

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	var countUserEmail int64
	database.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&countUserEmail)
	if countUserEmail > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Email already used",
		})
	}

	err := database.DB.Create(user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Created User",
		"data":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "No users present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Users Found",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	id := c.Params("userId")

	err := db.Find(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "No user present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "User Found",
		"data":    user,
	})
}
