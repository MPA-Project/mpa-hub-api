package oauth

import (
	"github.com/gofiber/fiber/v2"

	"myponyasia.com/hub-api/app/models"
	"myponyasia.com/hub-api/pkg/utils"
	"myponyasia.com/hub-api/platform/database"
)

// CreateUser func for creates a new user.
func CreateUser(c *fiber.Ctx) error {
	// Get now time.
	// now := time.Now().Unix()

	client := database.OpenDBConnection()

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Set initialized default data for book:
	// book.ID = uuid.New()
	// book.CreatedAt = time.Now()
	// book.BookStatus = 1 // 0 == draft, 1 == active

	// Validate book fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	hash, errHash := utils.HashPassword(user.Password)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errHash.Error(),
		})
	}

	err := client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(hash).
		Exec(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get claims from JWT.
	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	// Check, if received JSON data is valid.
	// if err := c.BodyParser(book); err != nil {
	// 	// Return status 400 and error message.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// Create database connection.
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	// Return status 500 and database connection error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// Create a new validator for a Book model.
	// validate := utils.NewValidator()

	// Set initialized default data for book:
	// book.ID = uuid.New()
	// book.CreatedAt = time.Now()
	// book.BookStatus = 1 // 0 == draft, 1 == active

	// Validate book fields.
	// if err := validate.Struct(book); err != nil {
	// 	// Return, if some fields are not valid.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// }

	// Delete book by given ID.
	// if err := db.CreateBook(book); err != nil {
	// 	// Return status 500 and error message.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
