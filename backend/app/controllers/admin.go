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

type Userdata struct {
	UserName string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.OpenDb()
	var users []models.User
	db.
		Where("deleted = ?", false).
		Find(&users)

	result, err := json.Marshal(&users)

	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func GetUser(c *fiber.Ctx) error {
	db := database.OpenDb()
	var user models.User
	db.First(&user, c.Params("id"))

	switch c.Params("action") {
	case "block":
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(400).JSON("Invalid ID")
		}
		if user.ID == uint(id) {
			user.Blocked = !user.Blocked
		}
	case "drop":
		pswd := os.Getenv("DEFAULT_PASSWORD")
		user.Password = utils.GeneratePassword(pswd)
	}
	db.Save(&user)

	result, err := json.Marshal(user)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(result)
}

func PostUser(c *fiber.Ctx) error {
	var userdata Userdata
	var user models.User

	if err := c.BodyParser(&userdata); err != nil {
		return c.Status(500).JSON(err)
	}

	db := database.OpenDb()
	db.
		Where("user_name LIKE ?", "%"+userdata.UserName+"%").
		First(&user)

	if user.ID == 0 {
		user.FullName = userdata.FullName
		user.UserName = userdata.UserName
		user.Email = userdata.Email
		pswd := os.Getenv("DEFAULT_PASSWORD")
		user.Password = utils.GeneratePassword(pswd)

	} else {
		return c.Status(400).JSON("User already exists")
	}
	db.Create(&user)
	return c.Status(201).JSON("User created")
}

func PatchUser(c *fiber.Ctx) error {
	var userdata Userdata
	var user models.User

	if err := c.BodyParser(&userdata); err != nil {
		return c.Status(500).JSON(err)
	}

	db := database.OpenDb()
	db.Where("user_name LIKE ?", "%"+userdata.UserName+"%").First(&user)

	if user.UserName == userdata.UserName {
		user.FullName = userdata.FullName
		user.UserName = userdata.UserName
		user.Email = userdata.Email
	}
	db.Save(&user)
	return c.Status(201).JSON("User updated")
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	db := database.OpenDb()
	db.First(&user, c.Params("id"))
	user.Deleted = true
	db.Save(&user)

	return c.Status(204).JSON("User deleted")
}

func GetRoles(c *fiber.Ctx) error {
	var user models.User
	var roleModel models.Role
	var hasRole bool

	db := database.OpenDb()
	db.
		Preload("Roles").
		First(&user, c.Params("user_id"))

	for _, role := range user.Roles {
		if role.NameRole == c.Params("value") {
			hasRole = true
			break
		}
	}
	if !hasRole {
		db.
			Where("name_role LIKE ?", "%"+c.Params("value")+"%").
			First(&roleModel)
		user.Roles = append(user.Roles, roleModel)
		db.Save(&user)
	}
	return c.Status(200).JSON("Role added")
}

func GetGroups(c *fiber.Ctx) error {
	var user models.User
	var groupModel models.Group
	var hasGroup bool

	db := database.OpenDb()
	db.
		Preload("Groups").
		First(&user, c.Params("user_id"))

	for _, group := range user.Groups {
		if group.NameGroup == c.Params("value") {
			hasGroup = true
			break
		}
	}
	if !hasGroup {
		db.Where("name_group LIKE ?", "%"+c.Params("value")+"%").First(&groupModel)
		user.Groups = append(user.Groups, groupModel)
		db.Save(&user)
	}
	return c.Status(200).JSON("Group added")
}

func DelRoles(c *fiber.Ctx) error {
	var user models.User
	var roleModel models.Role

	db := database.OpenDb()
	db.
		Preload("Roles").
		First(&user, c.Params("user_id"))

	for i, role := range user.Roles {
		if role.NameRole == c.Params("value") {
			db.Where("name_role LIKE ?", "%"+c.Params("value")+"%").First(&roleModel)
			user.Roles = append(user.Roles[:i], user.Roles[i+1:]...)
			db.Save(&user)
			break
		}
	}
	return c.Status(200).JSON("Role deleted")
}

func DelGroups(c *fiber.Ctx) error {
	var user models.User
	var groupModel models.Group

	db := database.OpenDb()
	db.
		Preload("Groups").
		First(&user, c.Params("user_id"))

	for i, group := range user.Groups {
		if group.NameGroup == c.Params("value") {
			db.Where("name_group LIKE ?", "%"+c.Params("value")+"%").First(&groupModel)
			user.Groups = append(user.Groups[:i], user.Groups[i+1:]...)
			db.Save(&user)
			break
		}
	}
	return c.Status(200).JSON("Group deleted")
}

func PostTablesRows(c *fiber.Ctx) error {
	var pagination = 16
	var hasPrev, hasNext bool

	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}

	searchData := struct {
		Search string `json:"search"`
	}{}
	err = c.BodyParser(&searchData)
	if err != nil {
		searchData.Search = ""
	}

	result := []interface{}{}

	db := database.OpenDb()
	query := db.
		Table(c.Params("item")).
		Limit(pagination).
		Offset(pagination * (intPage - 1)).
		Order("id desc")

	if searchData.Search != "" {
		query = query.Where("id = ?", searchData.Search)
	}
	query.Find(&result)

	if intPage > 1 {
		hasPrev = true
	}
	if len(result) == pagination {
		hasNext = true
	}
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}

func DelTableRows(c *fiber.Ctx) error {
	row := struct{}{}
	db := database.OpenDb()
	db.Table(c.Params("item"), c.Params("id")).First(row)
	db.Delete(row)
	db.Commit()

	return c.Status(204).JSON("Row deleted")
}
