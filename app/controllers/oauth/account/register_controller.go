package account

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
)

type SignupPayload struct {
	Username        string `json:"username" validate:"required,lte=255"`
	Email           string `json:"email" validate:"required,lte=255,email"`
	Password        string `json:"password" validate:"required,lte=25"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,lte=25,eqfield=password"`
	Token           string `json:"token" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	payload := new(SignupPayload)
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

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "signup"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var countUserEmail int64
	database.DB.Model(&models.User{}).Where("email = ?", payload.Email).Count(&countUserEmail)
	if countUserEmail > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Email already used",
		})
	}

	payload.Username = strings.ToLower(payload.Username)

	var countUserUsername int64
	database.DB.Model(&models.User{}).Where("username = ?", payload.Username).Count(&countUserUsername)
	if countUserUsername > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Username already used",
		})
	}

	var user = new(models.User)
	user.Email = payload.Email
	user.Username = payload.Username
	user.Password = payload.PasswordConfirm

	err := database.DB.Create(user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	token, err := utils.GenerateNewAccessToken(user.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Signup success",
		"data":    token,
	})
}
