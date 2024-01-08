package controllers

import (
	"encoding/json"
	"log"
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

	if user.UserName == userdata.UserName {
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

func GetRoles(c *fiber.Ctx) error {
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
	db.
		Preload("Roles").
		First(&user, c.Params("user_id"))

	for _, role := range user.Roles {
		if role.NameRole != c.Params("value") {
			var role models.Role
			db.Where("name_role LIKE ?", "%"+c.Params("value")+"%").First(&role)
			user.Roles = append(user.Roles, role)
			db.Save(&user)
			break
		}
	}
	db.First(&user, c.Params("user_id"))
	return c.Status(200).JSON(user)
}

func GetGroups(c *fiber.Ctx) error {
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
	db.
		Preload("Groups").
		First(&user, c.Params("user_id"))

	for _, group := range user.Groups {
		if group.NameGroup != c.Params("value") {
			var group models.Group
			db.Where("name_group LIKE ?", "%"+c.Params("value")+"%").First(&group)
			user.Groups = append(user.Groups, group)
			db.Save(&user)
			break
		}

	}
	db.First(&user, c.Params("user_id"))
	return c.Status(200).JSON(user)
}

func DelRoles(c *fiber.Ctx) error {
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

	db.
		Preload("Roles").
		First(&user, c.Params("user_id"))

	for i, role := range user.Roles {
		if role.NameRole == c.Params("value") {
			var role models.Role
			db.Where("name_role LIKE ?", "%"+c.Params("value")+"%").First(&role)
			user.Roles = append(user.Roles[:i], user.Roles[i+1:]...)
			db.Save(&user)
			break
		}
	}
	db.First(&user, c.Params("user_id"))
	return c.Status(200).JSON(user)
}

func DelGroups(c *fiber.Ctx) error {
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

	db.
		Preload("Groups").
		First(&user, c.Params("user_id"))

	for i, group := range user.Groups {
		if group.NameGroup == c.Params("value") {
			var group models.Group
			db.Where("name_group LIKE ?", "%"+c.Params("value")+"%").First(&group)
			user.Groups = append(user.Groups[:i], user.Groups[i+1:]...)
			db.Save(&user)
			break
		}
	}
	db.First(&user, c.Params("user_id"))
	return c.Status(200).JSON(user)
}

func PostTablesRows(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}
	var pagination = 16
	var hasPrev, hasNext bool
	items := []string{"user", "role", "group", "report", "resume", "connect"}

	searchData := struct {
		Search string `json:"search"`
	}{}
	if err := c.BodyParser(&searchData); err != nil {
		log.Println(err)
	}
	result := []interface{}{}

	for i := range items {
		if items[i] == c.Params("item") {
			query := db.
				Table(items[i]).
				Limit(pagination).
				Offset(pagination * (intPage - 1)).
				Order("id desc")

			if searchData.Search != "" {
				query = query.Where("id = ?", searchData.Search)
			}

			query.Find(&result)
		} else {
			query := db.
				Table(c.Params("item")).
				Limit(pagination).
				Offset(pagination * (intPage - 1)).
				Order("id desc")

			if searchData.Search != "" {
				query = query.Where("person_id = ?", searchData.Search)
			}

			query.Find(&result)
		}

		if intPage > 1 {
			hasPrev = true
		}
		if len(items[i]) == pagination {
			hasNext = true
		}
	}
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}

func DelTableRows(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"admin"}, []string{"admin"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	database.OpenDb().
		Table(c.Params("item")).
		Where("id = ?", c.Params("id")).
		Delete(c.Params("item"))

	return c.Status(204).JSON("Deleted")
}
