package account

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
	"myponyasia.com/hub-api/pkg/utils/authorization"
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

	recaptcha_disabled := os.Getenv("RECAPTCHA_DISABLED")
	if recaptcha_disabled != "true" {
		if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "signin"); err != nil && !resultCaptcha {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	var userEmail models.User
	err := database.DB.First(&userEmail, "email = ?", payload.Email).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Credentials inccorect",
		})
	}

	if !hash.CheckPasswordHash(payload.Password, userEmail.Password) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Credentials inccorect",
		})
	}

	at_token, rt_token, err := authorization.GenerateNewAccessToken(userEmail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	atExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRED"))
	rtExpired, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED"))

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Signin success",
		"data": fiber.Map{
			"access_token":          at_token,
			"acces_token_expired":   time.Now().Add(time.Duration(atExpired) * time.Minute),
			"refresh_token":         rt_token,
			"refresh_token_expired": time.Now().Add(time.Duration(rtExpired) * (24 * time.Hour)),
		},
	})
}
