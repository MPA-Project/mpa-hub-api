package uploads

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"myponyasia.com/hub-api/pkg/utils"
)

func UploadTemporary(c *fiber.Ctx) error {

	file, err := c.FormFile("file-upload")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "File upload not found",
		})
	}

	filename := uuid.New()

	c.SaveFile(file, fmt.Sprintf("./uploads/%s-%s", filename, file.Filename))

	info, err := os.Stat(fmt.Sprintf("./uploads/%s-%s", filename, file.Filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	go utils.RemoveExpiredFiles()

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":    false,
		"message":  "OK",
		"fileInfo": file,
		"info":     info,
	})
}
