package user

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"myponyasia.com/hub-api/pkg/database"
	"myponyasia.com/hub-api/pkg/entities"
	"myponyasia.com/hub-api/pkg/utils/universal"
	"myponyasia.com/hub-api/pkg/utils/uploads"
)

type SocialMediaList struct {
	Name string
	Url  string
}

type DonateList struct {
	Name string
	Url  string
}

func Me(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	uuid := claims["uuid"].(string)

	var user entities.User
	if err := database.DB.First(&user, "id = ?", uuid).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	var userImageAvatar *entities.FileManagerModelImage
	if err := json.Unmarshal([]byte(user.ProfilePicture), &userImageAvatar); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var avatarUrl string
	if userImageAvatar != nil {
		avatarUrl = universal.GenerateCDNUrl("user/profile", userImageAvatar)
		fmt.Fprintln(os.Stderr, avatarUrl)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "OK",
		"data": fiber.Map{
			"id":               user.ID.String(),
			"name":             user.Username,
			"email":            user.Email,
			"avatar":           avatarUrl,
			"avatarBackground": nil,
			"bio":              user.Bio,
			"socialMedia":      []SocialMediaList{},
			"donateLink":       []DonateList{},
		},
		"dbg": userImageAvatar,
	})
}

func UploadAvatar(c *fiber.Ctx) error {
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

	if file.Size >= (1 * 1024 * 1024) {
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
	var size bimg.ImageSize
	if size, err = readImage.Size(); size.Width > 500 || size.Height > 500 || err != nil {
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

	var user entities.User
	if err := database.DB.First(&user, "id = ?", userUuid).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	imageExtension := ".jpeg"
	constructFilename := fmt.Sprintf("%s%s", filename, imageExtension)
	modifiedFilename = fmt.Sprintf("./uploads/%s", constructFilename)
	constructPreviewUrl := fmt.Sprintf("%sutils/upload-temporary?filename=%s", os.Getenv("SERVER_URL_PUBLIC"), constructFilename)

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

	mtype := mimetype.Detect(newImage)
	year, month, day := time.Now().Date()

	// Insert to database
	var fileManager = new(entities.FileManager)
	fileManager.UserID = user.ID
	fileManager.Filename = filename
	fileManager.Extension = imageExtension
	fileManager.MimeType = mtype.String()
	fileManager.PYear = strconv.Itoa(year)
	fileManager.PMonth = fmt.Sprintf("%02d", month)
	fileManager.PDay = fmt.Sprintf("%02d", day)
	fileManager.Storage = "WASABI"
	fileManager.Filesize = int64(len(newImage))
	fileManager.UploadStatus = "UPLOADING"
	fileManager.ImageHeight = size.Height
	fileManager.ImageWidth = size.Width
	fileManager.ImageAvailableRes = []string{"original"}
	if err := database.DB.Create(&fileManager).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var userImageAvatar = new(entities.FileManagerModelImage)
	userImageAvatar.Filename = constructFilename
	userImageAvatar.ID = fileManager.ID
	userImageAvatar.PYear = fileManager.PYear
	userImageAvatar.PMonth = fileManager.PMonth
	userImageAvatar.PDay = fileManager.PDay
	userImageAvatar.ImageHeight = fileManager.ImageHeight
	userImageAvatar.ImageWidth = fileManager.ImageWidth
	userImageAvatar.ImageAvailableRes = fileManager.ImageAvailableRes
	userImageAvatarJson, err := json.Marshal(userImageAvatar)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	user.ProfilePicture = string(userImageAvatarJson)

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// S3 Upload
	uploadFilePath := fmt.Sprintf("user/profile/%s/%s/%s/%s", fileManager.PYear, fileManager.PMonth, fileManager.PDay, constructFilename)
	go uploads.UploadS3Update(newImage, uploadFilePath, *fileManager)

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
