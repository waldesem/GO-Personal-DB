package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"backend/app/models"
)

func GenerateNewAccessToken(user *models.User) (string, error) {
	secret, _ := os.LookupEnv("JWT_SECRET_KEY")
	exp, _ := os.LookupEnv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")
	minutesCount, _ := strconv.Atoi(exp)

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.NameRole)
	}

	var groups []string
	for _, group := range user.Groups {
		groups = append(groups, group.NameGroup)
	}

	// Create a new claims.
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.UserName,
		"fullname": user.FullName,
		"roles":    roles,
		"groups":   groups,
		"expires":  time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GenerateNewRefreshToken() (string, error) {
	secret, _ := os.LookupEnv("JWT_REFRESH_KEY")
	exp, _ := os.LookupEnv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT")
	minutesCount, _ := strconv.Atoi(exp)

	claims := jwt.MapClaims{}
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
