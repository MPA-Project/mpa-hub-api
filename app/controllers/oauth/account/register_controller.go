package account

import (
	"bytes"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/utils"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

type SignupPayload struct {
	Username        string `json:"username" validate:"required,lte=255"`
	Email           string `json:"email" validate:"required,lte=255,email"`
	Password        string `json:"password" validate:"required,lte=25"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,lte=25"`
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

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Passwords do not match",
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

	request_key := utils.RandomString(128, "alphanum") + "-" + hash.GetMD5Hash(user.ID.String()) + "-" + uuid.New().String()

	var user_request models.UserRequest
	if err := database.DB.Where("user_id = ?", user.ID).Where("request_type = ?", "EMAIL_VERIFICATION").Find(&user_request).Error; err != nil {
		user_request.UserID = user.ID
		user_request.RequestType = "EMAIL_VERIFICATION"
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

	tmpl := template.Must(template.ParseFiles("./views/email/request-email-verify.html"))
	email_data := TemplateEmail{
		Username: user.Username,
		Link:     os.Getenv("OAUTH_URL") + "?action=email-verification&redirect=signin&token=" + request_key + "&request=" + user.ID.String(),
	}
	tmpl_buffer := new(bytes.Buffer)
	tmpl.Execute(tmpl_buffer, email_data)
	if err := utils.SendHTML(user.Email, "Email Verification", tmpl_buffer.String()); err != nil {
		return c.Status(500).JSON(fiber.Map{
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
