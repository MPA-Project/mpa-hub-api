package account

import (
	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

type SigninPayload struct {
	Email    string `json:"email" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	Token    string `json:"token" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	payload := new(SigninPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "signin"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var userEmail models.User
	err := database.DB.Find(&userEmail, "email = ?", payload.Email).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Credentials inccorect",
			"data":    nil,
		})
	}

	if !hash.CheckPasswordHash(payload.Password, userEmail.Password) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Credentials inccorect",
			"data":    nil,
		})
	}

	token, err := utils.GenerateNewAccessToken(userEmail.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Signin success",
		"data":    token,
	})
}
