package middlewares

import (
	"os"

	"backend/pkg/utils"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)
}

// JWTProtected func for specify routes group with JWT authentication.
func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

func AuthRequired(roles []string, groups []string) func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	jwt := jwtMiddleware.New(config)

	// Return a function that combines the JWT middleware with custom authorization logic.
	return func(c *fiber.Ctx) error {
		// Use the JWT middleware to authenticate the request.
		err := jwt(c)

		// Check if there was an error or if the authentication failed.
		if err != nil || c.Locals("jwt") == nil {
			errMsg := "Unauthorized User"
			if err != nil {
				errMsg = err.Error()
			}
			return c.Status(401).JSON(errMsg)
		}

		// Custom authorization logic here.
		auth, err := utils.RolesGroupsInToken(c, roles, groups)
		if err != nil || auth == 0 {
			errMsg := "Unauthorized User"
			if err != nil {
				errMsg = err.Error()
			}
			return c.Status(401).JSON(errMsg)
		}

		// Proceed to the next middleware or handler.
		return c.Next()
	}
}
