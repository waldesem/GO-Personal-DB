package controllers

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/database"
)

func GetUsers(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	var users []models.User
	db.Find(&users)

	result, err := json.Marshal(users)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetUser(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	var user models.User
	db.First(&user, c.Params("id"))

	switch c.Params("action") {
	case "block":
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(400).JSON("Invalid ID")
		}
		if user.ID != uint(id) {
			user.Blocked = !user.Blocked
		}
	case "drop":
		user.Password = utils.GeneratePassword(os.Getenv("DEFAULT_PASSWORD"))
	}
	db.Save(&user)

	db.First(&user, c.Params("id"))
	result, err := json.Marshal(user)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func PostUser(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	userdata := struct {
		UserName string `json:"username"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}{}

	if err := c.BodyParser(&userdata); err != nil {
		return c.Status(500).JSON(err)
	}
	db := database.OpenDb()
	var user models.User
	db.Where("user_name LIKE ?", "%"+userdata.UserName+"%").First(&user)

	if user.UserName != userdata.UserName {
		user.FullName = userdata.FullName
		user.UserName = userdata.UserName
		user.Email = userdata.Email
		user.Password = utils.GeneratePassword(os.Getenv("DEFAULT_PASSWORD"))
	} else {
		return c.Status(400).JSON("User already exists")
	}
	db.Create(&user)
	return c.Status(201).JSON(user)
}

func PatchUser(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	userdata := struct {
		UserName string `json:"username"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}{}

	if err := c.BodyParser(&userdata); err != nil {
		return c.Status(500).JSON(err)
	}
	db := database.OpenDb()
	var user models.User
	db.Where("user_name LIKE ?", "%"+userdata.UserName+"%").First(&user)

	if user.UserName != userdata.UserName {
		user.FullName = userdata.FullName
		user.UserName = userdata.UserName
		user.Email = userdata.Email
		user.Password = utils.GeneratePassword(os.Getenv("DEFAULT_PASSWORD"))
	}
	db.Create(&user)
	return c.Status(201).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	var user models.User

	db.First(&user, c.Params("id"))
	db.Delete(&user)

	return c.Status(204).JSON("User deleted")
}
