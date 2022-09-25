package uploads

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"myponyasia.com/hub-api/pkg/utils"
)

func UploadTemporary(c *fiber.Ctx) error {

	go utils.RemoveExpiredFiles()

	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	userUuid := claims["uuid"].(string)

	file, err := c.FormFile("file-upload")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "File upload data not found",
		})
	}

	if file.Size >= (5 * 1024 * 1024) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "File upload to big",
		})
	}

	filename := fmt.Sprintf("%s--%s", userUuid, uuid.New())

	modifiedFilename := fmt.Sprintf("./uploads/temp-%s-%s", filename, file.Filename)

	c.SaveFile(file, modifiedFilename)

	buffer, err := bimg.Read(modifiedFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	defer os.Remove(modifiedFilename)

	readImage := bimg.NewImage(buffer)
	if size, err := readImage.Size(); size.Width >= 5000 || size.Height == 5000 || err != nil {
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

	constructFilename := fmt.Sprintf("%s%s", filename, ".jpeg")
	modifiedFilename = fmt.Sprintf("./uploads/%s", constructFilename)
	constructPreviewUrl := fmt.Sprintf("/utils/upload-temporary?filename=%s", constructFilename)

	newImage, err := readImage.Convert(bimg.JPEG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := bimg.Write(modifiedFilename, newImage); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// constructFilenameWebp := fmt.Sprintf("%s%s", filename, ".webp")
	// modifiedFilenameWebp := fmt.Sprintf("./uploads/%s", constructFilenameWebp)
	// newImage, err = readImage.Convert(bimg.WEBP)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": err.Error(),
	// 	})
	// }

	// if err := bimg.Write(modifiedFilenameWebp, newImage); err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
	// 		"error":   true,
	// 		"message": err.Error(),
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"filename": constructFilename,
			"path":     modifiedFilename,
			"preview":  constructPreviewUrl,
		},
	})
}

func UploadTemporaryViewer(c *fiber.Ctx) error {

	filename := c.Query("filename")

	if filename == "" || len(filename) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Filename is missing",
		})
	}

	modified_filename := fmt.Sprintf("./uploads/%s", filename)

	if _, err := os.Stat(modified_filename); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "File not found",
		})
	}

	return c.SendFile(modified_filename)
}
