package account

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/enums"
	"myponyasia.com/hub-api/pkg/utils"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

type TokenVerifyPayload struct {
	Action    string `json:"action" validate:"required,lte=255"`
	RequestID string `json:"requestId" validate:"required,lte=255"`
	Token     string `json:"token" validate:"required"`
	TokenKey  string `json:"tokenKey" validate:"required"`
}

type ForgotPasswordPayload struct {
	Email string `json:"email" validate:"required,lte=255"`
	Token string `json:"token" validate:"required"`
}

type ForgotPasswordConfirmPayload struct {
	RequestID       string `json:"requestId" validate:"required,lte=255"`
	Password        string `json:"password" validate:"required,lte=255"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,lte=255"`
	Token           string `json:"token" validate:"required"`
	TokenKey        string `json:"tokenKey" validate:"required"`
}

type EmailVerificationPayload struct {
	RequestID string `json:"requestId" validate:"required,lte=255"`
	Token     string `json:"token" validate:"required"`
	TokenKey  string `json:"tokenKey" validate:"required"`
}

type TemplateEmail struct {
	Username string
	Link     string
}

func tokenCheck(RequestID string, Action string, Token string) (models.UserRequest, error) {
	var token_hash = hash.GetMD5Hash(Token)
	var user_request models.UserRequest
	err := database.DB.Where("user_id = ?", RequestID).Where("request_type = ?", Action).Where("key_hash = ?", token_hash).First(&user_request).Error
	if err != nil {
		return models.UserRequest{}, err
	}

	if user_request.ExpiredAt.Before(time.Now()) {
		return models.UserRequest{}, errors.New("token expired")
	}

	return user_request, nil
}

func RequestTokenVerify(c *fiber.Ctx) error {
	payload := new(TokenVerifyPayload)
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

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "token_verify"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if !utils.EnumContains(enums.RequestEnum(), payload.Action) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Action not valid",
		})
	}

	if _, err := tokenCheck(payload.RequestID, payload.Action, payload.TokenKey); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Token valid",
	})
}

func ForgotPassword(c *fiber.Ctx) error {
	payload := new(ForgotPasswordPayload)
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

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "forgot_password"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var user models.User
	if err := database.DB.First(&user, "email = ?", payload.Email).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Account not found",
		})
	}

	fmt.Println("=========")
	fmt.Println(user)
	fmt.Println("=========")

	request_key := utils.RandomString(128, "alphanum") + "-" + hash.GetMD5Hash(user.ID.String()) + "-" + uuid.New().String()

	var user_request models.UserRequest
	if err := database.DB.Where("user_id = ?", user.ID).Where("request_type = ?", "RESET_PASSWORD").First(&user_request).Error; err != nil {
		user_request.UserID = user.ID
		user_request.RequestType = "RESET_PASSWORD"
		user_request.Key = request_key
		user_request.KeyHash = hash.GetMD5Hash(request_key)
		user_request.ExpiredAt = time.Now().Add(time.Hour * 2)

		if err := database.DB.Create(user_request).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	} else {
		user_request.Key = request_key
		user_request.KeyHash = hash.GetMD5Hash(request_key)
		user_request.ExpiredAt = time.Now().Add(time.Hour * 2)

		if err := database.DB.Save(user_request).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	tmpl := template.Must(template.ParseFiles("./views/email/request-reset-password.html"))
	email_data := TemplateEmail{
		Username: user.Username,
		Link:     os.Getenv("OAUTH_URL") + "?action=reset-password&redirect=signin&token=" + request_key + "&request=" + user.ID.String(),
	}
	tmpl_buffer := new(bytes.Buffer)
	tmpl.Execute(tmpl_buffer, email_data)
	if err := utils.SendHTML(user.Email, "Forgot Password", tmpl_buffer.String()); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Forgot password success",
	})
}

func ForgotPasswordConfirm(c *fiber.Ctx) error {
	payload := new(ForgotPasswordConfirmPayload)
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

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Passwords do not match",
		})
	}

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "forgot_password_confirm"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	user_request, err := tokenCheck(payload.RequestID, "RESET_PASSWORD", payload.TokenKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", user_request.UserID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	hash, errHash := hash.HashPassword(payload.Password)
	if errHash != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errHash.Error(),
		})
	}
	user.Password = hash

	if err := database.DB.Save(user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := database.DB.Delete(&user_request).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Password changed",
	})
}

func EmailVerification(c *fiber.Ctx) error {
	payload := new(EmailVerificationPayload)
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

	if resultCaptcha, err := utils.CaptchaVerifyToken(payload.Token, "email_verification"); err != nil && !resultCaptcha {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	user_request, err := tokenCheck(payload.RequestID, "EMAIL_VERIFICATION", payload.TokenKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", user_request.UserID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}
	user.EmailVerify = true
	user.EmailVerifyAt = time.Now()

	if err := database.DB.Save(user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := database.DB.Delete(&user_request).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Email verified",
	})
}
