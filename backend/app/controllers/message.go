package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/database"
)

func GetMessages(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"user"}, []string{"staffsec"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	intPage, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		intPage = 1
	}

	db := database.OpenDb()
	var messages []models.Message
	var pagination = 16
	var hasPrev, hasNext bool

	switch c.Params("action") {
	case "new":
		db.
			Where("status = ?, user_id = ?", "new", auth).
			Find(&messages).
			Limit(pagination).
			Offset(pagination * (intPage - 1))
	case "read":
		db.
			Where("status = ?, user_id = ?", "new", auth).
			Find(&messages)
		for message := range messages {
			messages[message].StatusRead = "reply"
			db.Save(&messages[message])
		}
		db.
			Where("status = ?, user_id = ?", "new", auth).
			Find(&messages).
			Limit(pagination).
			Offset(pagination * (intPage - 1))
	}

	if intPage > 1 {
		hasPrev = true
	}
	if len(messages) == pagination {
		hasNext = true
	}
	result, err := json.Marshal(messages)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(fiber.Map{"result": result, "hasNext": hasNext, "hasPrev": hasPrev})
}

func DeleteMessage(c *fiber.Ctx) error {
	auth, err := utils.RolesGroupsInToken(c, []string{"user"}, []string{"staffsec"})
	if err != nil || auth == 0 {
		errMsg := "Unauthorized User"
		if err != nil {
			errMsg = err.Error()
		}
		return c.Status(401).JSON(errMsg)
	}

	db := database.OpenDb()
	var messages []models.Message

	db.
		Find(&messages).
		Where("person_id = ?", auth).
		Delete(&messages)
	return c.SendStatus(fiber.StatusOK)
}
