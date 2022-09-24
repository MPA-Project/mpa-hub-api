package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"myponyasia.com/hub-api/pkg/utils/authorization"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTSessionProtected() func(*fiber.Ctx) error {

	key, err := jwt.ParseRSAPublicKeyFromPEM(authorization.PublicKey)
	if err != nil {
		return nil
	}

	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningMethod:  "RS256",
		SigningKey:     key,
		ContextKey:     "jwt",
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSessionSuccess,
	}

	return jwtMiddleware.New(config)
}

func JWTRefreshProtected() func(*fiber.Ctx) error {

	key, err := jwt.ParseRSAPublicKeyFromPEM(authorization.PublicKey)
	if err != nil {
		return nil
	}

	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningMethod:  "RS256",
		SigningKey:     key,
		ContextKey:     "jwt",
		ErrorHandler:   jwtError,
		SuccessHandler: jwtRefreshSuccess,
	}

	return jwtMiddleware.New(config)
}

func jwtSessionSuccess(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	if typ := claims["typ"].(string); typ != "session" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid token type",
		})
	}

	return c.Next()
}

func jwtRefreshSuccess(c *fiber.Ctx) error {
	jwtClaims := c.Locals("jwt").(*jwt.Token)
	claims := jwtClaims.Claims.(jwt.MapClaims)

	if typ := claims["typ"].(string); typ != "refresh" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid token type",
		})
	}

	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
