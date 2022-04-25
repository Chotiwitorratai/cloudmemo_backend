package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)


func JWTProtected() func(*fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET_KEY")
	config := jwtMiddleware.Config{
		SigningKey:   []byte(secret),
		ContextKey:   "jwt", 
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "error",
		"message":   err.Error(),
	})
}

