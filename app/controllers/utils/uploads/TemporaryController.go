package uploads

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"myponyasia.com/hub-api/pkg/utils"
)

func UploadTemporary(c *fiber.Ctx) error {

	go utils.RemoveExpiredFiles()

	file, err := c.FormFile("file-upload")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "File upload not found",
		})
	}

	if file.Size >= (5 * 1024 * 1024) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "File upload to big",
		})
	}

	filename := uuid.New()

	modified_filename := fmt.Sprintf("./uploads/temp-%s-%s", filename, file.Filename)

	c.SaveFile(file, modified_filename)

	buffer, err := bimg.Read(modified_filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	defer os.Remove(modified_filename)

	read_image := bimg.NewImage(buffer)
	if size, err := read_image.Size(); size.Width >= 5000 || size.Height == 5000 || err != nil {
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": "Image to large",
		})
	}

	newImage, err := read_image.Convert(bimg.JPEG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	modified_filename = fmt.Sprintf("./uploads/%s%s", filename, ".jpeg")
	if err := bimg.Write(modified_filename, newImage); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"filename": modified_filename,
		},
	})
}
