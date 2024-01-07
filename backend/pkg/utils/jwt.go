package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"backend/platform/database"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

func GenerateNewTokens(id uint, roles []string, groups []string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := GenerateNewAccessToken(id, roles, groups)
	if err != nil {
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := GenerateNewRefreshToken(id)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func GenerateNewAccessToken(id uint, roles []string, groups []string) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	fmt.Println(minutesCount)

	// Create a new claims.
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["roles"] = roles
	claims["groups"] = groups
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GenerateNewRefreshToken(id uint) (string, error) {
	secret := os.Getenv("JWT_REFRESH_KEY")
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID  uint
	Roles   []string
	Groups  []string
	Expires int64
}

// ApplyToken checks if the provided roles and groups are present in the token metadata.
func RolesGroupsInToken(c *fiber.Ctx, roles []string, groups []string) (uint, error) {
	tokenMeta, err := ExtractTokenMetadata(c)
	if err != nil {
		return 0, err
	}
	// Set expiration time from JWT data of current book.
	expires := tokenMeta.Expires
	if time.Now().Unix() > expires {
		fmt.Println(time.Now().Unix())
		fmt.Println(expires)
		return 0, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	hasGroup := parseRolesGroups(groups, tokenMeta.Groups)
	hasRole := parseRolesGroups(roles, tokenMeta.Roles)

	if hasGroup && hasRole {
		return tokenMeta.UserID, nil
	} else {
		return 0, nil
	}
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

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	bearToken := c.Get("Authorization")

	_, err := database.RedisConnection().Get(c.Context(), bearToken).Result()
	if err != redis.Nil {
		return nil, err
	}

	onlyToken := strings.Split(bearToken, " ")
	token, err := jwt.Parse(onlyToken[1], jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	// Setting and checking token
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["id"].(float64)
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
			UserID:  uint(userUint),
			Roles:   rolesSlice,
			Groups:  groupsSlice,
			Expires: expires,
		}, nil
	}
	return nil, err
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
