package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/middlewares"
	"backend/pkg/utils"
	"backend/platform/cache"
	"backend/platform/database"
)

type UserData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	NewPswd  string `json:"new_pswd"`
}

type Tokens struct {
	Access  string
	Refresh string
}

func GetLogin(c *fiber.Ctx) error {
	tokenMeta, _ := middlewares.ExtractTokenMetadata(c)
	db := database.OpenDb()
	var user models.User
	db.
		Preload("Roles").
		Preload("Groups").
		First(&user, tokenMeta.UserID)

	result, err := json.Marshal(user)
	fmt.Println(string(result))
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func PostLogin(c *fiber.Ctx) error {
	var userdata UserData
	if err := c.BodyParser(&userdata); err != nil {
		log.Println(err)
	}

	db := database.OpenDb()
	var user models.User

	db.
		Preload("Roles").
		Preload("Groups").
		Where("user_name LIKE ?", "%"+userdata.UserName+"%").
		First(&user)

	if user.ID != 0 && !user.Blocked {
		if utils.ComparePasswords(user.Password, userdata.Password) {
			deltaChange := time.Since(user.CreatedAt)

			if !user.UpdatedAt.IsZero() && deltaChange.Hours() < 365*24 {
				user.LastLogin = time.Now()
				user.Attempt = 0
				db.Save(&user)

				tokens := Tokens{}
				var err error
				tokens.Access, err = utils.GenerateNewAccessToken(&user)
				tokens.Refresh, _ = utils.GenerateNewRefreshToken()

				if err != nil {
					return c.Status(200).JSON("Denied")
				}
				result := map[string]interface{}{
					"message": "Authenticated",
					"tokens":  tokens,
				}
				return c.Status(200).JSON(result)
			}
			return c.Status(200).JSON("Expired")
		} else {
			if user.Attempt < 9 {
				user.Attempt++
			} else {
				user.Blocked = true
			}
			db.Save(&user)
		}
	}
	return c.Status(401).JSON("Denied")
}

func PatchLogin(c *fiber.Ctx) error {
	var userdata UserData
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
				return c.Status(200).JSON("Denied")
			}
			db.Save(&user)
			return c.Status(201).JSON("Authenticated")
		}
	}
	return c.Status(200).JSON("Denied")
}

// DeleteLogin deletes the login for a given user.
func DeleteLogin(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	err := cache.RedisConnection().Set(c.Context(), token, true, time.Hour).Err()
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func RefreshToken(c *fiber.Ctx) error {
	// Get refresh token from storage.
	claims, err := middlewares.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	db := database.OpenDb()
	var user models.User
	db.
		Preload("Roles").
		Preload("Groups").
		First(&user, claims.UserID)

	if user.ID != 0 && !user.Blocked {
		accessToken, err := utils.GenerateNewAccessToken(&user)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(accessToken)
	}
	return c.Status(401).JSON("unauthorized")
}
