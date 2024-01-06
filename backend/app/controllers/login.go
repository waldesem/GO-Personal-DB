package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/database"
)

func GetLogin(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if time.Now().Unix() > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	db := database.OpenDb()
	var user models.User
	db.First(&user, claims.UserID)

	result, err := json.Marshal(user)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func PostLogin(c *fiber.Ctx) error {
	userdata := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&userdata); err != nil {
		log.Println(err)
	}

	db := database.OpenDb()
	var user models.User
	var result map[string]interface{}

	db.
		Where("user_name LIKE ?", "%"+userdata.UserName+"%").
		First(&user)

	if user.ID != 0 && !user.Blocked {
		if utils.ComparePasswords(user.Password, userdata.Password) {
			deltaChange := time.Since(user.CreatedAt)

			if !user.UpdatedAt.IsZero() && deltaChange.Hours() < 365*24 {
				user.LastLogin = time.Now()
				user.Attempt = 0
				db.Save(&user)

				tokens, err := utils.GenerateNewTokens(string(rune(user.ID)))
				if err != nil {
					result = map[string]interface{}{
						"message": "Denied",
					}
				}
				result := map[string]interface{}{
					"message": "Authenticated",
					"tokens":  tokens,
				}

				return c.Status(200).JSON(result)
			}
			result := map[string]interface{}{
				"message": "Expired",
			}
			return c.Status(200).JSON(result)
		} else {
			if user.Attempt < 9 {
				user.Attempt++
			} else {
				user.Blocked = true
			}
			db.Save(&user)
		}
	}
	result = map[string]interface{}{
		"message": "Denied",
	}
	return c.Status(401).JSON(result)
}

func PatchLogin(c *fiber.Ctx) error {
	userdata := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		NewPswd  string `json:"new_pswd"`
	}{}
	if err := c.BodyParser(&userdata); err != nil {
		log.Println(err)
	}

	db := database.OpenDb()
	var user models.User
	db.
		Where("user_name LIKE ?", "%"+userdata.UserName+"%").
		First(&user)

	if user.ID != 0 && !user.Blocked {
		if utils.ComparePasswords(user.Password, userdata.Password) {
			user.Password = utils.GeneratePassword(userdata.NewPswd)
			if user.Password == nil {
				result := map[string]interface{}{
					"message": "Denied",
				}
				return c.Status(200).JSON(result)
			}
			db.Save(&user)
			result := map[string]interface{}{
				"message": "Authenticated",
			}
			return c.Status(201).JSON(result)
		}
	}
	result := map[string]interface{}{
		"message": "Denied",
	}
	return c.Status(200).JSON(result)
}

func DeleteLogin(c *fiber.Ctx) error {
	// Delete token from storage.
	c.Locals("access_token", nil)
	c.Locals("refresh_token", nil)
	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	// Get refresh token from storage.
	refreshToken, err := utils.GenerateNewRefreshToken()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(refreshToken)
}
