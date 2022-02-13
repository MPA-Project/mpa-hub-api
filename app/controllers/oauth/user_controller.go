package oauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Validate custom validation
	validate := utils.NewValidator()
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	var userEmail []models.User
	database.DB.Find(&userEmail, "email = ?", user.Email)
	if len(userEmail) > 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "Email has been used",
			"data":    nil,
		})
	}

	// Create the Note and return error if encountered
	err := database.DB.Create(user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return the created note
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Created User",
		"data":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	// find all users in the database
	db.Find(&users)

	// If no users is present return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "No users present",
			"data":    nil,
		})
	}

	// Else return users
	return c.JSON(fiber.Map{
		"error":   true,
		"message": "Users Found",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "No user present",
			"data":    nil,
		})
	}

	// Return the user with the Id
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "User Found",
		"data":    user,
	})
}
