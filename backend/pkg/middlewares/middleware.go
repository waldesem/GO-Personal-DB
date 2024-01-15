package middlewares

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"backend/platform/cache"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID   uint
	FullName string
	UserName string
	Roles    []string
	Groups   []string
	Expires  int64
}

// FiberMiddleware provide Fiber's built-in middlewares.
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)
}

func AuthRequired(roles []string, groups []string) func(*fiber.Ctx) error {
	secret, _ := os.LookupEnv("JWT_SECRET_KEY")
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(secret)},
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
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		tokenMeta, err := ExtractTokenMetadata(c)
		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		expires := tokenMeta.Expires
		if time.Now().Unix() > expires {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "expired",
			})
		}

		hasGroup := parseRolesGroups(groups, tokenMeta.Groups)
		hasRole := parseRolesGroups(roles, tokenMeta.Roles)

		if hasGroup && hasRole {
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   "denied",
			})
		}
	}
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	bearToken := c.Get("Authorization")

	cache.RedisConnection().Get(c.Context(), bearToken).Result()

	onlyToken := strings.Split(bearToken, " ")
	token, err := jwt.Parse(onlyToken[1], jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	// Setting and checking token
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["id"].(float64)
		userName := claims["fullname"].(string)
		userLogin := claims["username"].(string)

		userUint, err := strconv.ParseUint(fmt.Sprintf("%.0f", userID), 10, 64)
		if err != nil {
			return nil, err
		}
		expires := int64(claims["expires"].(float64))

		rolesSlice := []string{} // Initialize empty string slice
		if rolesInterface, ok := claims["roles"].([]interface{}); ok {
			for _, role := range rolesInterface {
				if roleStr, ok := role.(string); ok {
					rolesSlice = append(rolesSlice, roleStr)
				}
			}
		}
		groupsSlice := []string{} // Initialize empty string slice
		if groupsInterface, ok := claims["groups"].([]interface{}); ok {
			for _, group := range groupsInterface {
				if groupStr, ok := group.(string); ok {
					groupsSlice = append(groupsSlice, groupStr)
				}
			}
		}
		return &TokenMetadata{
			UserID:   uint(userUint),
			FullName: userName,
			UserName: userLogin,
			Roles:    rolesSlice,
			Groups:   groupsSlice,
			Expires:  expires,
		}, nil
	}
	return nil, err
}

// parseRolesGroups is a function that takes in two slices of strings, values and metas, and returns a boolean value.
func parseRolesGroups(values []string, metas []string) bool {
	if len(values) == 0 {
		return true
	}
	for _, value := range values {
		for _, meta := range metas {
			if value == meta {
				return true
			}
		}
	}
	return false
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

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	secret, _ := os.LookupEnv("JWT_SECRET_KEY")
	return []byte(secret), nil
}
