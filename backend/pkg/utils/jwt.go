package utils

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

func GenerateNewTokens(id string, roles []string, groups []string) (*Tokens, error) {
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

func GenerateNewAccessToken(id string, roles []string, groups []string) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")
	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

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

func GenerateNewRefreshToken(id string) (string, error) {
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
	UserID  uuid.UUID
	Roles   []string
	Groups  []string
	Expires int64
}

// ApplyToken checks if the provided roles and groups are present in the token metadata.
func ApplyToken(c *fiber.Ctx, roles []string, groups []string) (uuid.UUID, error) {
	TokenMetadata, err := ExtractTokenMetadata(c)
	if err != nil {
		return uuid.Nil, err
	}
	// Set expiration time from JWT data of current book.
	expires := TokenMetadata.Expires
	if time.Now().Unix() > expires {
		return uuid.Nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	hasGroup := false
	hasRole := false

	if len(TokenMetadata.Groups) == 0 {
		hasGroup = true
	} else {
		for _, role := range roles {
			for _, r := range TokenMetadata.Roles {
				if r == role {
					hasRole = true
					break
				}
			}
		}
	}

	if len(TokenMetadata.Groups) == 0 {
		hasGroup = true
	} else {
		for _, group := range groups {
			for _, g := range TokenMetadata.Groups {
				if g == group {
					hasGroup = true
					break
				}
			}
		}
	}

	if hasGroup && hasRole {
		return TokenMetadata.UserID, nil
	} else {
		return uuid.Nil, nil
	}
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}
		roles, err := json.Marshal(claims["roles"])
		if err != nil {
			return nil, err
		}
		groups, err := json.Marshal(claims["groups"])
		if err != nil {
			return nil, err
		}

		// Expires time.
		expires := int64(claims["expires"].(float64))

		return &TokenMetadata{
			UserID:  userID,
			Roles:   strings.Split(string(roles), ","),
			Groups:  strings.Split(string(groups), ","),
			Expires: expires,
		}, nil
	}

	return nil, err
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
