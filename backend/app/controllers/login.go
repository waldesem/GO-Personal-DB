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
	userId, err := utils.ApplyToken(c, []string{}, []string{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db := database.OpenDb()
	var user models.User
	db.First(&user, userId)

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

				var roles []string
				for _, role := range user.Roles {
					roles = append(roles, role.NameRole)
				}

				var groups []string
				for _, group := range user.Groups {
					groups = append(groups, group.NameGroup)
				}

				tokens, err := utils.GenerateNewTokens(string(rune(user.ID)), roles, groups)
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
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db := database.OpenDb()
	var user models.User
	db.First(&user, claims.UserID)

	var roles []string
	var groups []string

	if user.ID != 0 && !user.Blocked {
		for _, role := range user.Roles {
			roles = append(roles, role.NameRole)
		}

		for _, group := range user.Groups {
			groups = append(groups, group.NameGroup)
		}

		refreshToken, err := utils.GenerateNewAccessToken(
			claims.UserID.String(), roles, groups,
		)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(refreshToken)
	}
	return c.Status(401).JSON("unauthorized")
}
